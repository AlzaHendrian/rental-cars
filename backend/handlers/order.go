package handler

import (
	ordersdto "backend/dto/orders"
	dto "backend/dto/results"
	"backend/models"
	"backend/repositories"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
	CarRepository   repositories.CarRepository
	UserRepository  repositories.UserRepository
}

func HandlerOrders(OrderRepository repositories.OrderRepository, CarRepository repositories.CarRepository, UserRepository repositories.UserRepository) *handlerOrder {
	return &handlerOrder{OrderRepository, CarRepository, UserRepository}
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
	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		fmt.Println(err, "<<< errror")
		return err
	} else {
		fmt.Printf("%v", my_data)
	}

	fmt.Printf("%v", my_data)
	userLogin := c.Get("userLogin")

	fmt.Println(userLogin, "<<< user Login")
	userID := int(userLogin.(jwt.MapClaims)["id"].(float64))
	fmt.Println(userID, "<<< user IDnya")
	// carID, _ := strconv.Atoi(c.FormValue("car_id"))
	pickupDateStr := my_data["pickup_date"].(string)
	timeDropoffDateStr := my_data["dropoff_date"].(string)
	pickupDate, err := time.Parse("2006-01-02", pickupDateStr)
	if err != nil {
		fmt.Println(pickupDate, "<<< err pickup date")
	}
	timeDropoffDate, err := time.Parse("2006-01-02", timeDropoffDateStr)
	if err != nil {
		fmt.Println(timeDropoffDate, "<<< err TimeDropoff date")
	}

	fmt.Println("SEBELUM REQUEST")

	request := ordersdto.CreateRequestOrder{
		UserID:          userID,
		CarID:           int(my_data["car_id"].(float64)),
		PickupDate:      pickupDate,
		TimeDropoffDate: timeDropoffDate,
		PickupLocation:  my_data["pickup_location"].(string),
		DropoffLocation: my_data["dropoff_location"].(string),
	}

	fmt.Println(request, "<<< request order")

	car, err := h.CarRepository.GetCar(request.CarID)
	if err != nil {
		fmt.Println(err, "<<< ini error car request")
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		fmt.Println(err, "Error Validation")
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	order := models.Order{
		UserID:          request.UserID,
		CarID:           car.CarID,
		Car:             car,
		PickupDate:      request.PickupDate,
		TimeDropoffDate: request.TimeDropoffDate,
		PickupLocation:  request.PickupLocation,
		DropoffLocation: request.DropoffLocation,
	}

	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		fmt.Println(err, "Error data creating order")
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
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
		UserID:          u.UserID,
		CarID:           u.CarID,
		PickupDate:      u.PickupDate,
		TimeDropoffDate: u.TimeDropoffDate,
		PickupLocation:  u.PickupLocation,
		DropoffLocation: u.DropoffLocation,
	}
}
