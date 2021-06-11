package app

import (
	"encoding/json"
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ar *dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		ar.CustomerId = vars["customer_id"]
		account, err := a.service.NewAccount(ar)

		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
