package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"server.com/pkg/utils"

	"gorm.io/gorm"
)

var (
	gormDb *gorm.DB
	err    error
)

// MysqlConnection func for connection to Mysql database.
func MysqlConnection() (*gorm.DB, error) {
	if gormDb == nil {
		gormDb, err = NewMysqlConnection()
	}

	return gormDb, err
}

func NewMysqlConnection() (*gorm.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build Mysql connection URL.
	mysqlConnURL, _ := utils.ConnectionURLBuilder("mysql")

	// Define database connection for Mysql.
	db, err := gorm.Open(mysql.Open(mysqlConnURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get db\n")
	}
	// Set database connection settings:
	//   - SetMaxOpenConns: the default is 0 (unlimited)
	//   - SetMaxIdleConns: defaultMaxIdleConns = 2
	//   - SetConnMaxLifetime: 0, connections are reused forever
	sql.SetMaxOpenConns(maxConn)
	sql.SetMaxIdleConns(maxIdleConn)
	sql.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	return db, nil
}
