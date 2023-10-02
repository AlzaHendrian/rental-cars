package routes

import (
	handler "backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoute(e *echo.Group) {
	OrderRepository := repositories.RepositoryOrder(mysql.DB)
	UserRepository := repositories.RepositoryUser(mysql.DB)
	CarRepository := repositories.RepositoryCar(mysql.DB)
	h := handler.HandlerOrders(OrderRepository, UserRepository, CarRepository)

	e.GET("/orders", h.FindOrders)
	e.GET("/order/:id", h.GetOrder)
	e.POST("/order", middleware.Auth(h.CreateOrder))
	e.DELETE("/order/:id", middleware.Auth(h.DeleteOrder))
}
