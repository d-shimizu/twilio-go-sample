package migrate

import (
	"fmt"
	"localhost/twilio-go-sample/infra/database/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.PhoneNumber{}); err != nil {
		return fmt.Errorf("failed to migrate PhoneNumber model: %w", err)
	}

	return nil
}
