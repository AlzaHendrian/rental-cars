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
	h := handler.HandlerOrders(OrderRepository)

	e.GET("/orders", h.FindOrders)
	e.GET("/order/:id", h.GetOrder)
	e.POST("/order", middleware.Auth(middleware.UploadFile(h.CreateOrder)))
	e.PATCH("/order/:id", middleware.Auth(middleware.UploadFile(h.UpdateOrder)))
	e.DELETE("/order/:id", middleware.Auth(h.DeleteOrder))
}
