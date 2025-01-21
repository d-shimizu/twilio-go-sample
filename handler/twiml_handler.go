package handler

import (
	"localhost/twilio-go-sample/domain/model"
	"localhost/twilio-go-sample/usecase"
	"net/http"
)

type TwiMLHandler struct {
	useCase *usecase.TwiMLUseCase
}

func NewTwiMLHandler(useCase *usecase.TwiMLUseCase) *TwiMLHandler {
	return &TwiMLHandler{
		useCase: useCase,
	}
}

func (h *TwiMLHandler) HandleVoice(w http.ResponseWriter, r *http.Request) {
	req := &model.VoiceRequest{
		From: r.FormValue("From"),
		To:   r.FormValue("To"),
	}

	twiml, err := h.useCase.HandleIncomingCall(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(twiml))
}
