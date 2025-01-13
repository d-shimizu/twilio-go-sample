package handler

import (
	"encoding/json"
	"localhost/twilio-go-sample/usecase"
	"net/http"
)

type PhoneNumberHandler struct {
	phoneNumberUseCase usecase.PhoneNumberUseCase
}

func NewPhoneNumberHandler(phoneNumberUseCase usecase.PhoneNumberUseCase) *PhoneNumberHandler {
	return &PhoneNumberHandler{
		phoneNumberUseCase: phoneNumberUseCase,
	}
}

type purchasePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

func (h *PhoneNumberHandler) PurchasePhoneNumber(w http.ResponseWriter, r *http.Request) {
	var req purchasePhoneNumberRequest
	var ctx = r.Context()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	purchasedNumber, err := h.phoneNumberUseCase.PurchasePhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchasedNumber)
}
