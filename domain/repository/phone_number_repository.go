package repository

import (
	"context"
	"localhost/twilio-go-sample/domain/model"
)

type PhoneNumberRepository interface {
	Create(ctx context.Context, phoneNumber *model.PhoneNumber) error
	//FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.PhoneNumber, error)
}
