package main

import (
	"localhost/twilio-go-sample/handler"
	"localhost/twilio-go-sample/infra/twilio"
	"localhost/twilio-go-sample/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Twilioクライアントの初期化
	twilioClient := twilio.NewClient(accountSid, authToken)

	// ユースケースの初期化
	phoneNumberUseCase := usecase.NewPhoneNumberUseCase(twilioClient)

	// ハンドラーの初期化
	phoneNumberHandler := handler.NewPhoneNumberHandler(*phoneNumberUseCase)

	// ルーターの設定
	r := mux.NewRouter()
	r.HandleFunc("/phone-numbers/purchase", phoneNumberHandler.PurchasePhoneNumber).Methods(http.MethodPost)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":18080", r))
}
