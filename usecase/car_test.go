package usecase


import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "testing"
    "time"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

var (
    sqlEngine   = &mock.Mock{}
    CarsTable   = &mock.Mock{}
    requestService = usecase.NewRequestService(sqlEngine, CarsTable)
    carRow = &domain.CarRequest{
        RegistrationNumber: "test",
        Model:             "test",
        Mileage:           1000,
    }
)
func AddCarNominalTest(t *testing.T) {
    requestService.On("FindCarByRegistrationNumber", mock.Anything).Return(nil, nil)
    requestService.carTable.On("Insert", mock.Anything, mock.Anything).Return(nil)
    err := requestService.AddCar(context.Background(), carRow)
    require.Nil(t, err)
}

func AddCarExistTest(t *testing.T) {
    requestService.On("FindCarByRegistrationNumber", mock.Anything).Return( &domain.Car{
		RegistrationNumber: "test",
		Model:             "test",
		Mileage:           1000,
		Rented:            false,
	}, nil)
    err := requestService.AddCar(context.Background(), carRow)
    require.NotNil(t, err)
    require.Equal(t, "cannot add this car,It existed", err.Error())
	require.Nil(t, err)
}

func RentCarTest(t *testing.T) {
    requestService.On("FindCarByRegistrationNumber", mock.Anything).Return( &domain.Car{
		RegistrationNumber: "test",
		Model:             "test",
		Mileage:           1000,
		Rented:            false,
	}, nil)
	requestService.carTable.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
    err := requestService.RentCar(context.Background(), carRow)
	require.Nil(t, err)
}

func RentCarNomTest(t *testing.T) {
    requestService.On("FindCarByRegistrationNumber", mock.Anything).Return( &domain.Car{
		RegistrationNumber: "test",
		Model:             "test",
		Mileage:           1000,
		Rented:            true,
	}, nil)
	requestService.carTable.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
    err := requestService.ReturnCar(context.Background(), carRow)
	require.Nil(t, err)
}