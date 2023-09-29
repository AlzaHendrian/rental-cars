package routes

import (
	handler "backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/labstack/echo/v4"
)

func CarsRoute(e *echo.Group) {
	CarRepository := repositories.RepositoryCar(mysql.DB)
	h := handler.HandlerCars(CarRepository)

	e.GET("/cars", h.FindCar)
	e.GET("/car/:id", h.GetCar)
	e.POST("/car", middleware.UploadFile(h.CreateCar))
	e.PATCH("/car/:id", middleware.UploadFile(h.UpdateCar))
	e.DELETE("/car/:id", h.DeleteCar)
}
