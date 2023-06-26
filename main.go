package main

import (
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/customers"
	"github.com/edwintrumpet/prueba-tecnica-t-evolvers/internal/server"
	"go.uber.org/fx"
)

// @title           enerBit API
// @version         0.1
// @description     API to manage customers and work orders
// @host      localhost:3000
// @BasePath  /
func main() {
	fx.New(fx.Options(
		fx.Provide(customers.NewController),
		fx.Provide(customers.NewServices),
		fx.Invoke(server.Start),
	))
}
