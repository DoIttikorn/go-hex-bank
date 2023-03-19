package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Doittikorn/bank/service"
)

type customerHandler struct {
	customerSrv service.CustomerService
}

func NewCustomerHandler(customerSrv service.CustomerService) *customerHandler {
	return &customerHandler{customerSrv: customerSrv}
}

func (h *customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customer, err := h.customerSrv.GetCustomer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// return customer response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// convert customer to []byte
	json.NewEncoder(w).Encode(customer)
}
