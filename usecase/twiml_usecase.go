package usecase

import "localhost/twilio-go-sample/domain/model"

type TwiMLUseCase struct{}

func NewTwiMLUseCase() *TwiMLUseCase {
	return &TwiMLUseCase{}
}

func (uc *TwiMLUseCase) HandleIncomingCall(req *model.VoiceRequest) (string, error) {
	return `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say language="ja-JP">お電話ありがとうございます。これはテスト応答です。</Say>
</Response>`, nil
}
