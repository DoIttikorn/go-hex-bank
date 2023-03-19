package main

import (
	"database/sql"
	"net/http"

	"github.com/Doittikorn/bank/handler"
	"github.com/Doittikorn/bank/repository"
	"github.com/Doittikorn/bank/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		println(err.Error())
	}
	defer db.Close()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomer).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)

}
