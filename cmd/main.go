package main

import (
	"fmt"
	"localhost/twilio-go-sample/handler"
	"localhost/twilio-go-sample/infra/database/repository"
	"localhost/twilio-go-sample/infra/twilio"
	"localhost/twilio-go-sample/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// データベース接続
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Twilioクライアントの初期化
	// インフラ層の初期化
	twilioClient := twilio.NewClient(accountSid, authToken)
	phoneNumberRepo := repository.NewPhoneNumberRepository(db)

	// ユースケースの初期化
	phoneNumberUseCase := usecase.NewPhoneNumberUseCase(twilioClient, phoneNumberRepo)

	// ハンドラーの初期化
	phoneNumberHandler := handler.NewPhoneNumberHandler(*phoneNumberUseCase)

	// ユースケースの初期化
	twimlUseCase := usecase.NewTwiMLUseCase()

	// ハンドラーの初期化
	twimlHandler := handler.NewTwiMLHandler(twimlUseCase)

	// ルーターの設定
	r := mux.NewRouter()
	r.HandleFunc("/phone-numbers/purchase", phoneNumberHandler.PurchasePhoneNumber).Methods(http.MethodPost)
	r.HandleFunc("/phone-numbers/available", phoneNumberHandler.ListAvailablePhoneNumber).Methods(http.MethodGet)
	r.HandleFunc("/twiml/voice", twimlHandler.HandleVoice).Methods(http.MethodPost)

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}

	// サーバーの起動
	//log.Fatal(http.ListenAndServe(":18080", r))
	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
