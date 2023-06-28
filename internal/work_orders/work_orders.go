package workorders

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ansel1/merry"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type service struct {
	db  *gorm.DB
	rdb *redis.Client
}

type Service interface {
	Create(models.CreateWorkOrder) (*models.WorkOrder, error)
	Finish(req models.FinishWorkOrder) (*models.WorkOrder, error)
	ListAll(req models.ListAllWorkOrders) ([]models.WorkOrder, error)
}

func New(db *gorm.DB, rdb *redis.Client) Service {
	return &service{
		db:  db,
		rdb: rdb,
	}
}

func (s *service) Create(c models.CreateWorkOrder) (*models.WorkOrder, error) {
	/* -------------------------- Data validation -------------------------- */
	if err := c.Validate(); err != nil {
		return nil, merry.Wrap(err).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(err.Error())
	}

	if c.PlannedDateBegin.After(c.PlannedDateEnd) {
		errMsg := "plannedDateEnd must be after plannedDateBegin"
		return nil, merry.New(errMsg).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(errMsg)
	}

	if c.PlannedDateBegin.Add(time.Hour * 2).Before(c.PlannedDateEnd) {
		errMsg := "time between plannedDates must be a max of two hours"
		return nil, merry.New(errMsg).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(errMsg)
	}

	/* -------------------------- Type conversion -------------------------- */
	// Convert a CreateWorkerOrder into a WorkerOrder to save it
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	data := models.WorkOrder{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, merry.Wrap(err)
	}

	data.Status = models.New

	/* -------------------------- Save on database -------------------------- */
	tx := s.db.Begin()
	tx.Create(&data)
	/* ------------------------ Update the customer ------------------------ */
	tx.Model(&models.Customer{}).
		Where("id = ?", c.CustomerID).
		Updates(map[string]interface{}{
			"is_active": false,
			"end_date":  time.Now().UTC(),
		}).Scan(&data.Customer)
	res := tx.Commit()
	if err := res.Error; err != nil {
		return nil, merry.Wrap(err)
	}

	return &data, nil
}

func (s *service) Finish(req models.FinishWorkOrder) (*models.WorkOrder, error) {
	/* -------------------------- Data validation -------------------------- */
	if err := req.Validate(); err != nil {
		return nil, merry.Wrap(err).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(err.Error())
	}

	workOrder := models.WorkOrder{}
	/* -------------------------- Update customer -------------------------- */
	tx := s.db.Begin()
	tx.Model(&workOrder).
		Where("id = ?", req.WorkOrderID).
		Update("status", models.Done).
		Scan(&workOrder)
	/* -------------------------- Update customer -------------------------- */
	tx.Model(&workOrder.Customer).
		Where("id = ?", req.CustomerID).
		Updates(map[string]interface{}{
			"is_active":  true,
			"start_date": time.Now().UTC(),
		}).Scan(&workOrder.Customer)

	/* ------------------------ Send event to Redis ------------------------ */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "finished-orders",
		ID:     "*",
		Values: map[string]string{
			"WorkOrderId": req.WorkOrderID,
			"customerId":  req.CustomerID,
			"status":      string(workOrder.Status),
		},
	}).Err(); err != nil {
		return nil, merry.Wrap(err)
	}

	res := tx.Commit()
	if err := res.Error; err != nil {
		return nil, merry.Wrap(err)
	}

	return &workOrder, nil
}

func (s *service) ListAll(req models.ListAllWorkOrders) ([]models.WorkOrder, error) {
	/* -------------------------- Data validation -------------------------- */
	if err := req.Validate(); err != nil {
		return nil, merry.Wrap(err).
			WithHTTPCode(http.StatusBadRequest).
			WithUserMessage(err.Error())
	}

	tx := s.db

	if req.Since != "" {
		since, err := time.Parse(time.RFC3339, req.Since)
		if err != nil {
			return nil, merry.Wrap(err).
				WithHTTPCode(http.StatusBadRequest).
				WithUserMessage("since format should be in iso date")
		}

		tx = tx.Where("created_at >= ?", since)
	}

	if req.Until != "" {
		until, err := time.Parse(time.RFC3339, req.Until)
		if err != nil {
			return nil, merry.Wrap(err).
				WithHTTPCode(http.StatusBadRequest).
				WithUserMessage("until format should be in iso date")
		}

		tx = tx.Where("created_at <= ?", until)
	}

	if req.Status != "" {
		tx = tx.Where("status = ?", req.Status)
	}

	workOrders := []models.WorkOrder{}

	res := tx.Preload(clause.Associations).Find(&workOrders)
	if err := res.Error; err != nil {
		return nil, merry.Wrap(err)
	}

	return workOrders, nil
}
