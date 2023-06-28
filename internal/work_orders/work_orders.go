package workorders

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ansel1/merry"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/models"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	Create(models.CreateWorkOrder) (*models.WorkOrder, error)
	Finish(req models.FinishWorkOrder) (*models.WorkOrder, error)
}

func New(db *gorm.DB) Service {
	return &service{
		db: db,
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
	// stream

	res := tx.Commit()
	if err := res.Error; err != nil {
		return nil, merry.Wrap(err)
	}

	return &workOrder, nil
}
