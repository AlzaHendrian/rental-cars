package handler

import (
	carsdto "backend/dto/cars"
	dto "backend/dto/results"
	"backend/models"
	"backend/repositories"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handler struct {
	CarRepository repositories.CarRepository
}

func HandlerCars(CarRepository repositories.CarRepository) *handler {
	return &handler{CarRepository}
}

func (h *handler) FindCar(c echo.Context) error {
	cars, err := h.CarRepository.FindCars()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cars})
}

func (h *handler) GetCar(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	car, err := h.CarRepository.GetCar(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: car})
}

func (h *handler) CreateCar(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	dayRate, _ := strconv.Atoi(c.FormValue("day_rate"))
	monthRate, _ := strconv.Atoi(c.FormValue("month_rate"))

	request := carsdto.CreateCar{
		CarName:   c.FormValue("car_name"),
		DayRate:   float64(dayRate),
		MonthRate: float64(monthRate),
		Image:     dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	fmt.Println(API_KEY, CLOUD_NAME, API_SECRET)

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "cars"})

	if err != nil {
		fmt.Println(err.Error())
	}

	car := models.Cars{
		CarName:   request.CarName,
		DayRate:   request.DayRate,
		MonthRate: request.MonthRate,
		Image:     resp.SecureURL,
	}

	data, err := h.CarRepository.CreateCar(car)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) UpdateCar(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	dayRate, _ := strconv.Atoi(c.FormValue("day_rate"))
	monthRate, _ := strconv.Atoi(c.FormValue("month_rate"))

	request := carsdto.CreateCar{
		CarName:   c.FormValue("car_name"),
		DayRate:   float64(dayRate),
		MonthRate: float64(monthRate),
		Image:     dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	fmt.Println(API_KEY, CLOUD_NAME, API_SECRET)

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "cars"})

	id, _ := strconv.Atoi(c.Param("id"))

	car, err := h.CarRepository.GetCar(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	car.CarName = request.CarName
	car.DayRate = request.DayRate
	car.MonthRate = request.MonthRate
	car.Image = resp.SecureURL

	data, err := h.CarRepository.UpdateCar(car)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) DeleteCar(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	car, err := h.CarRepository.GetCar(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CarRepository.DeleteCar(car, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func convertResponse(u models.Cars) carsdto.ResponseCars {
	return carsdto.ResponseCars{
		CarName:   u.CarName,
		DayRate:   u.DayRate,
		MonthRate: u.MonthRate,
		Image:     u.Image,
	}
}
