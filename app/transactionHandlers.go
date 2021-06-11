package app

import (
	"encoding/json"
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (t TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var tr *dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		tr.AccountId = vars["account_id"]
		transaction, err := t.service.NewTransaction(tr)

		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusOK, transaction)
		}
	}
}
