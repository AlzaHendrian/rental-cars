package handler

import (
	ordersdto "backend/dto/orders"
	dto "backend/dto/results"
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrders(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) FindOrders(c echo.Context) error {
	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: orders})
}

func (h *handlerOrder) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: order})
}

func (h *handlerOrder) CreateOrder(c echo.Context) error {
	pickupDateStr := c.FormValue("pickup_date")
	timeDropoffDateStr := c.FormValue("dropoff_date")
	pickupDate, _ := time.Parse("2006-01-02", pickupDateStr)
	timeDropoffDate, _ := time.Parse("2006-01-02", timeDropoffDateStr)

	request := ordersdto.CreateRequestOrder{
		PickupDate:      pickupDate,
		TimeDropoffDate: timeDropoffDate,
		PickupLocation:  c.FormValue("pickup_location"),
		DropoffLocation: c.FormValue("dropoff_location"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	order := models.Order{
		PickupDate:      request.PickupDate,
		TimeDropoffDate: request.TimeDropoffDate,
		PickupLocation:  request.PickupLocation,
		DropoffLocation: request.DropoffLocation,
	}

	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseOrder(data)})
}

func (h *handlerOrder) UpdateOrder(c echo.Context) error {
	pickupDateStr := c.FormValue("pickup_date")
	timeDropoffDateStr := c.FormValue("dropoff_date")
	pickupDate, _ := time.Parse("2006-01-02", pickupDateStr)
	timeDropoffDate, _ := time.Parse("2006-01-02", timeDropoffDateStr)

	request := ordersdto.CreateRequestOrder{
		PickupDate:      pickupDate,
		TimeDropoffDate: timeDropoffDate,
		PickupLocation:  c.FormValue("pickup_location"),
		DropoffLocation: c.FormValue("dropoff_location"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	order.PickupDate = request.PickupDate
	order.TimeDropoffDate = request.TimeDropoffDate
	order.PickupLocation = request.PickupLocation
	order.DropoffLocation = request.DropoffLocation

	data, err := h.OrderRepository.UpdateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseOrder(data)})
}

func (h *handlerOrder) DeleteOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.OrderRepository.DeleteOrder(order, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseOrder(data)})
}

func convertResponseOrder(u models.Order) ordersdto.ResponseOrder {
	return ordersdto.ResponseOrder{
		PickupDate:      u.PickupDate,
		TimeDropoffDate: u.TimeDropoffDate,
		PickupLocation:  u.PickupLocation,
		DropoffLocation: u.DropoffLocation,
	}
}
