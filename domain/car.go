package domain

type Car struct {
	Model              string
	RegistrationNumber string
	Mileage            int64
	Rented             bool
}

type CarRequest struct {
	Model              string
	RegistrationNumber string
	Mileage            int64
}

type CarReturnRequest struct {
	Mileage            int64
}
