package main

import (
	"net/http"
	"time"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/eya20/cars/controllers"
	"github.com/eya20/cars/domain/usecase"
	"github.com/gorilla/mux"
)

var (
	requestController controllers.Request
)

func main() {
	// DATABASE
	database := CreateDatabase()
	dbEngine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			database,
			information_schema.NewInformationSchemaDatabase(),
		))

	requestService := usecase.NewRequestService(dbEngine, CarsTable)
	requestController = controllers.NewRequestController(requestService)
	router := routes()
	httpServer := &http.Server{
		Addr:              "3000",
		WriteTimeout:      time.Second * 15,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 2,
		IdleTimeout:       time.Second * 60,
		Handler:           router,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Panicf("listen: %s", err)
	}

}
func routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/cars", requestController.Add).Methods("POST")
	router.HandleFunc("/cars", requestController.GetAllCars).Methods("GET")
	return router
}
