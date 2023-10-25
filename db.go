package main

import (
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/sql"
)

const (
	DbName        = "test-car"
	CarsTableName = "cars"
)

var (
	CarsTable *memory.Table
)

func CreateDatabase() *memory.Database {
	db := memory.NewDatabase(DbName)
	CarsTable = memory.NewTable(CarsTableName, sql.Schema{
		{Name: "registrationNumber", Type: sql.Text, Nullable: false, Source: CarsTableName},
		{Name: "model", Type: sql.Text, Nullable: false, Source: CarsTableName},
		{Name: "mileage", Type: sql.Int32, Nullable: false, Source: CarsTableName},
		{Name: "rented", Type: sql.Boolean, Nullable: false, Source: CarsTableName},
		[]sql.Index{
			sql.Index{
				Name:    "unique_registrationNumber",
				Unique:  true,
				Columns: []string{"registrationNumber"},
			},
		},
	})
	db.AddTable(CarsTableName, CarsTable)
	return db
}
