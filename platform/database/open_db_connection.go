package database

import (
	"os"

	repositories "github.com/xiaoma/trading/app/repository"
	"gorm.io/gorm"
)

// Queries struct for collect all app queries.
type DB struct {
	*gorm.DB
	Forex *repositories.Trading // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*DB, error) {
	// Define Database connection variables.
	var (
		db  *gorm.DB
		err error
	)

	// Get DB_TYPE value from .env file.
	dbType := "mysql"

	if dbTypeEnv := os.Getenv("DB_TYPE"); dbTypeEnv != "" {
		dbType = dbTypeEnv
	}

	// Define a new Database connection with right DB type.
	switch dbType {
	case "mysql":
		db, err = MysqlConnection()
	}

	if err != nil {
		return nil, err
	}

	return &DB{
		DB:    db,
		Forex: repositories.TradingRepo(db), // add self-create-repo
	}, nil
}
