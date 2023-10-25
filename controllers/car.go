package controllers

import (
	"bytes"
	"context"
	"net/http"

	"github.com/eya20/cars/domain"
	"github.com/eya20/cars/domain/response"
	"github.com/eya20/cars/usecase"
)

type Request interface {
	Add(rw http.ResponseWriter, req *http.Request)
}

type request struct {
	requestService usecase.Request
}

func NewRequestController(s usecase.Request) Request {
	return request{requestService: s}
}

// @Summary Add new car
// @Description Add a new car and return car status/error
// @Tags car
// @Accept json
// @Param request body type car
// @Produce json
// @Success 202 {object}
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 500 {object} response.Error "InternalServerError"
// @Router /cars [POST]
func (r request) Add(rw http.ResponseWriter, req *http.Request) {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(req.Body)
	defer req.Body.Close()
	if err != nil {
		err = response.Wrap(err, response.ErrBodyRead, "cannot read body")
		return
	}
	carRequest := new(domain.CarRequest)
	err = &r.requestService.AddCar(context.Background(), carRequest)
	if err != nil {
		err = response.Wrap(err, "Mysql error", "cannot add the car")
		return
	}
	w.WriteHeader(http.StatusOK)
    message := []byte("Car added successfully")
    _, _ = rw.Write(message)
}

// @Summary Get all cars
// @Description get all cars return aray of car /error
// @Tags car
// @Accept json
// @Param request body type car
// @Produce json
// @Success 202 {object}
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 500 {object} response.Error "InternalServerError"
// @Router /cars [GET]
func (r request) GetAllCars(rw http.ResponseWriter, req *http.Request) {
	cars, err = &r.requestService.FindAllCars(context.Background())
	if err != nil {
		err = response.Wrap(err, "Mysql error", "can not get Cars")
		return
	}

	 jsonResponse, err := json.Marshal(cars)
	 if err != nil {
		 http.Error(rw, err.Error(), http.StatusInternalServerError)
		 return
	 }
 	 rw.Header().Set("Content-Type", "application/json")
 	 _, _ = rw.Write(jsonResponse)
}

// @Summary Rent a car
// @Description rent a car return status/error
// @Tags car
// @Accept json
// @Produce json
// @Success 202 {object}
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 500 {object} response.Error "InternalServerError"
// @Router  /cars/:registration/rentals [POST]
func (r request) RentCar(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
    registration := vars["registration"]
	err = &r.requestService.RentCar(context.Background(),registration)
	if err != nil {
		err = response.Wrap(err, "Error :", "can not Rent the car")
		return
	}

	w.WriteHeader(http.StatusOK)
    message := []byte("Car rented successfully")
    _, _ = rw.Write(message)
}

// @Summary Rent a car
// @Description rent a car return status/error
// @Tags car
// @Accept json
// @Produce json
// @Success 202 {object}
// @Failure 400 {object} response.Error "Bad Request"
// @Failure 500 {object} response.Error "InternalServerError"
// @Router  /cars/:registration/returns [POST]
func (r request) ReturnCar(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
    registration := vars["registration"]
	var requestBody domain.CarReturnRequest
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	err = &r.requestService.ReturnCar(context.Background(),registration,requestBody.Mileage)
	if err != nil {
		err = response.Wrap(err, "Error :", "can not return the car")
		return
	}

	w.WriteHeader(http.StatusOK)
    message := []byte("Car returned successfully")
    _, _ = rw.Write(message)
}