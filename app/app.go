package app

import (
	"github.com/alibaihaqi/banking/domain"
	"github.com/alibaihaqi/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub()) }
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)
	//router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// Starting Server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
