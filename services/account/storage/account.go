package storage

import (
	"github.com/zale144/nanosapp/services/account/model"
	"github.com/zale144/nanosapp/services/account/db"
)

type AccountStorage struct {}

// GetByUsername retrieves an account from the database matching the provided username
func (as AccountStorage) GetByUsername(username string) (*model.Account, error) {

	var account model.Account
	tx := db.PgsqlDB.Where(model.Account{Username: username})

	err := tx.First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Insert adds an account row to the database
func (as AccountStorage) Insert(account model.Account) error {

	tx := db.PgsqlDB.Begin()

	if err := tx.Save(&account).Error; err != nil {
		if db.IsUniqueConstraintError(err, model.UniqueConstraintUsername) {
			return &model.UsernameDuplicateError{}
		}
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		if db.IsUniqueConstraintError(err, model.UniqueConstraintUsername) {
			return &model.UsernameDuplicateError{}
		}
		tx.Rollback()
		return err
	}
	return nil
}