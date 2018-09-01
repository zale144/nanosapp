package db


import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/zale144/nanosapp/services/account/model"
	"github.com/lib/pq"
)

var (
	DBInfo  string
	PgsqlDB *gorm.DB
)

// openDB starts the connection with the database
func openDB() (*gorm.DB, error) {
	if PgsqlDB != nil {
		return PgsqlDB, nil
	}
	var err error
	PgsqlDB, err = gorm.Open("postgres", DBInfo)
	if err != nil {
		return nil, err
	}
	PgsqlDB.DB().SetMaxOpenConns(400)
	PgsqlDB.DB().SetMaxIdleConns(0)
	PgsqlDB.DB().SetConnMaxLifetime(100 * time.Second)
	return PgsqlDB, nil
}

// InitDB initializes the database and migrates the data
func InitDB() error {
	_, err := openDB()
	if err != nil {
		return err
	}

	// Add tables here
	PgsqlDB.AutoMigrate(&model.Account{})

	// Enable Logger, show detailed log
	PgsqlDB.LogMode(true)
	return nil
}

// IsUniqueConstraintError returns true if the type of error is unique constraint violation
func IsUniqueConstraintError(err error, constraintName string) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" && pqErr.Constraint == constraintName
	}
	return false
}