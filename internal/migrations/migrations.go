package migrations

import (
	orgModels "jna-manager/internal/domain/models/org"
	paymentModels "jna-manager/internal/domain/models/payments"
	userModels "jna-manager/internal/domain/models/users"
)

func GetMigrationModels() []interface{} {
	mgrModel := []interface{}{
		&userModels.User{},
		&orgModels.Blog{},
		&orgModels.Beneficiary{},
		&paymentModels.Donation{},
		&paymentModels.Payment{},
	}

	return mgrModel
}
