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

func (uc *PhoneNumberUseCase) ListAvailablePhoneNumbers(ctx context.Context, numberType, areaCode string) ([]*model.AvailablePhoneNumber, error) {
	var numbers []*model.AvailablePhoneNumber

	switch numberType {
	case "local":
		resp, err := uc.twilioClient.ListAvailableLocalPhoneNumbers(areaCode)
		if err != nil {
			return nil, fmt.Errorf("failed to get local numbers: %w", err)
		}

		for _, num := range *resp {
			if num.PhoneNumber == nil {
				continue
			}
			number := &model.AvailablePhoneNumber{
				PhoneNumber:  *num.PhoneNumber,
				FriendlyName: *num.FriendlyName,
			}
			if num.Locality != nil {
				number.Locality = *num.Locality
			}
			if num.Region != nil {
				number.Region = *num.Region
			}
			if num.PostalCode != nil {
				number.PostalCode = *num.PostalCode
			}
			number.Capabilities.Voice = num.Capabilities.Voice
			number.Capabilities.SMS = num.Capabilities.Sms
			number.Capabilities.MMS = num.Capabilities.Mms

			numbers = append(numbers, number)
		}

	case "toll-free":
		resp, err := uc.twilioClient.ListAvailableTollFreePhoneNumbers()
		if err != nil {
			return nil, fmt.Errorf("failed to get toll-free numbers: %w", err)
		}

		for _, num := range *resp {
			if num.PhoneNumber == nil {
				continue
			}
			number := &model.AvailablePhoneNumber{
				PhoneNumber:  *num.PhoneNumber,
				FriendlyName: *num.FriendlyName,
			}
			number.Capabilities.Voice = num.Capabilities.Voice
			number.Capabilities.SMS = num.Capabilities.Sms
			number.Capabilities.MMS = num.Capabilities.Mms

			numbers = append(numbers, number)
		}

		//	case "mobile":
		//		resp, err := uc.twilioClient.ListAvailableMobileNumbers(areaCode)
		//		if err != nil {
		//			return nil, fmt.Errorf("failed to get mobile numbers: %w", err)
		//		}
		//
		//		for _, num := range resp.Mobile {
		//			if num.PhoneNumber == nil {
		//				continue
		//			}
		//			number := &model.AvailablePhoneNumber{
		//				PhoneNumber:  *num.PhoneNumber,
		//				FriendlyName: *num.FriendlyName,
		//			}
		//			if num.Region != nil {
		//				number.Region = *num.Region
		//			}
		//			number.Capabilities.Voice = *num.Capabilities.Voice
		//			number.Capabilities.SMS = *num.Capabilities.Sms
		//			number.Capabilities.MMS = *num.Capabilities.Mms
		//
		//			numbers = append(numbers, number)
		//		}

	default:
		return nil, fmt.Errorf("invalid number type: %s", numberType)
	}

	return numbers, nil
}
