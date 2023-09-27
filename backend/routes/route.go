package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	CarsRoute(e)
	UserRoutes(e)
	AuthRoutes(e)
	OrderRoute(e)
}
