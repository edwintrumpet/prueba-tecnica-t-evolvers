package main

import (
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/api"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/config"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/db"
	redisservice "github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/redis_service"
	workorders "github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/work_orders"
	"go.uber.org/fx"
)

// @title           enerBit API
// @version         0.1
// @description     API to manage customers and work orders
// @host      localhost:3000
// @BasePath  /
func main() {
	fx.New(fx.Options(
		fx.Provide(config.New),
		fx.Provide(db.NewDB),
		fx.Provide(redisservice.New),
		fx.Provide(customers.New),
		fx.Provide(workorders.New),
		fx.Provide(api.NewServer),
		fx.Invoke(api.Server.Start),
	)).Run()
}
