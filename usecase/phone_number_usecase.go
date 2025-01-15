package usecase

import (
	"context"
	"fmt"
	"localhost/twilio-go-sample/domain/model"
	"localhost/twilio-go-sample/domain/repository"
	ucrepo "localhost/twilio-go-sample/usecase/repository"
)

type PhoneNumberUseCase struct {
	twilioClient    ucrepo.PhoneNumberRepository
	phoneNumberRepo repository.PhoneNumberRepository
}

func NewPhoneNumberUseCase(client ucrepo.PhoneNumberRepository, repo repository.PhoneNumberRepository) *PhoneNumberUseCase {
	return &PhoneNumberUseCase{
		twilioClient:    client,
		phoneNumberRepo: repo,
	}
}

func (uc *PhoneNumberUseCase) PurchasePhoneNumber(ctx context.Context, phoneNumber string) (*model.PhoneNumber, error) {
	resp, err := uc.twilioClient.PurchasePhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	if resp.PhoneNumber == nil || resp.AccountSid == nil {
		return nil, fmt.Errorf("invalid response from Twilio: missing required fields")
	}

	purchasePhoneNumber := &model.PhoneNumber{
		AccountSid:  *resp.AccountSid,
		PhoneNumber: *resp.PhoneNumber,
	}

	if err := uc.phoneNumberRepo.Create(ctx, purchasePhoneNumber); err != nil {
		return nil, fmt.Errorf("failed to save phone number: %w", err)
	}

	return purchasePhoneNumber, nil
}
