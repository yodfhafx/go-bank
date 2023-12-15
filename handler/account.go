package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yodfhafx/go-bank/errs"
	"github.com/yodfhafx/go-bank/service"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("Request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	response, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
