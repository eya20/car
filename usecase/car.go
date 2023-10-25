package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/sql"

	"github.com/eya20/cars/domain"
)

type Request interface {
	AddCar(ctx context.Context, carRequest domain.CarRequest) error
	FindCarByRegistrationNumber(ctx context.context,registrationNumber string) (error, *domain.Car)
	ReturnCar(ctx context.context,registrationNumber string, mileage int64) (error)
	RentCar(ctx context.context,registrationNumber string) (error)
	FindAllCars(ctx context.context) (error, []domain.Car)
}

type request struct {
	sqlEngine *sqle.Engine
	carTable  *memory.Table
}

func NewRequestService(sqlEngine *sqle.Engine, carTable *memory.Table) Request {
	return request{
		sqlEngine: sqlEngine,
		carTable:  carTable,
	}
}

func (r request) AddCar(ctx context.Context, carRequest domain.CarRequest) error {
	sqlContext := sql.NewContext(ctx)
	row , err = r.FindCarByRegistrationNumber(ctx,carRequest.RegistrationNumber)
	if row != nil {
		return errors.New(fmt.Sprintf("cannot add this car , It existed "))
	} else {
		rowCar := sql.NewRow(
			carRequest.registrationNumber,
			carRequest.Model,
			carRequest.Mileage,
			false,
		)
		err = r.carTable.Insert(sqlContext, rowCar)
		if err != nil {
			return errors.New(fmt.Sprintf("cannot save row: %v", err))
		}
		return nil
	}

}
 func (r request) FindCarByRegistrationNumber(ctx context, carRegistrationNumber string) (error , *domain.Car){
	sqlContext := sql.NewContext(ctx)
	query := fmt.Sprintf("SELECT * FROM `%s` WHERE `registrationNumber` = '%s' ", r.carTable.Name(), carRegistrationNumber)
	log.Print(query)

	_, iter, err := r.sqlEngine.Query(sql.NewEmptyContext().WithCurrentDB(boot.DbName), query)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot execute request: %v", err))
	}
	row, err := iter.Next()
	if err != nil {
		return errors.New(fmt.Sprintf("error in searching car : ", err)), nil 
	}
	var car domain.Car 
	if err := row.To(&car); err != nil { 
		return errors.New(fmt.Sprintf("can not map data in car struct : %v", err)), result
 	}
	return car, nil
 }
 func (r request) FindAllCars(ctx context.context) (error, []domain.Car) {
	sqlContext := sql.NewContext(ctx)
	query := fmt.Sprintf("SELECT * FROM `%s` ", r.carTable.Name())
	log.Print(query)

	_, iter, err := r.sqlEngine.Query(sql.NewEmptyContext().WithCurrentDB(boot.DbName), query)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot execute request: %v", err))
	}
	row, err := iter.Next()
	if err != nil {
		return errors.New(fmt.Sprintf("error in searching car : ", err)), nil 
	}

	// Create an array to store the results 
	var results []domain.Car // Iterate over the result and create struct instances 
	for { 
		row, err := iter.Next() 
		if err == sql.ErrNoRows { 
		break 
		// No more rows 
		} else if err != nil { 
			return errors.New(fmt.Sprintf("Error :", err)), result
		} 
			var car domain.Car 
			if err := row.To(&car); err != nil { 
				return errors.New(fmt.Sprintf("can not map data in car struct : %v", err)), result
 			}
			 // Append the struct instance to the results array 
			 results = append(results, car) 
		} 
	return nil , result
 }

func (r request) RentCar(ctx context.context,registrationNumber string) (error) {
    car , err := r.FindCarByRegistrationNumber(ctx, registrationNumber)
	if err != nil {
		return errors.New(fmt.Sprintf("Error in searching car : %v", err))
	}
	if car == nil {
		return return errors.New(fmt.Sprintf("car doesn't exist : %v", err))
	}
	if car.rented {
		return return errors.New(fmt.Sprintf("car already rented :"))
	}
	_, err = r.carTable.Update(ctx, sql.NewRow("rented", rented), "registration_number = ?", registrationNumber)

    if err != nil {
        return errors.New(fmt.Sprintf("error updating the 'rented' field: %v", err))
    }
	return nil
}

func (r request) ReturnCar(ctx context.context,registrationNumber string, mileage int64) (error) {
	car , err := r.FindCarByRegistrationNumber(ctx, registrationNumber)
	if err != nil {
		return errors.New(fmt.Sprintf("Error in searching car : %v", err))
	}
	if car == nil {
		return return errors.New(fmt.Sprintf("car doesn't exist : %v", err))
	}
	if !car.rented {
		return return errors.New(fmt.Sprintf("car was not market as rented"))
	}
	car.Mileage += 	mileage
	_, err = r.carTable.Update(ctx, sql.NewRow("rented", false,"mileage",car.Mileage), "registration_number = ?", registrationNumber)

    if err != nil {
        return errors.New(fmt.Sprintf("error updating the 'rented' field: %v", err))
    }
	return nil 
}
