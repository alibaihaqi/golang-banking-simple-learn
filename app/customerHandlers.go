package app

import (
	"encoding/json"
	"fmt"
	"github.com/alibaihaqi/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(p)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessgae())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer, err := ch.service.GetCustomer(vars["customer_id"])
	fmt.Println(vars["customer_id"], customer)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessgae())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
