package repository

import (
	"context"
	"localhost/twilio-go-sample/domain/model"
	dbmodel "localhost/twilio-go-sample/infra/database/model"

	"gorm.io/gorm"
)

type phoneNumberRepository struct {
	db *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) *phoneNumberRepository {
	return &phoneNumberRepository{db: db}
}

func (r *phoneNumberRepository) Create(ctx context.Context, phoneNumber *model.PhoneNumber) error {
	dbPhoneNumber := &dbmodel.PhoneNumber{
		TwilioSid:   phoneNumber.AccountSid, // 注: 実際のTwilio SIDに応じて調整が必要
		PhoneNumber: phoneNumber.PhoneNumber,
		AccountSid:  phoneNumber.AccountSid,
		Status:      "active",
	}

	return r.db.WithContext(ctx).Create(dbPhoneNumber).Error
}
