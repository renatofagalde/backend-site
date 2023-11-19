package mysqldb

import (
	"backend-site/src/config/logger"
	"backend-site/src/model/repository/entity"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"os"

	"gorm.io/gorm"
)

var (
	MYSQL_DB_URL      = "MYSQL_DB_URL"
	MYSQL_DB_PORT     = "MYSQL_DB_PORT"
	MYSQL_DB_SCHEMA   = "MYSQL_DB_SCHEMA"
	MYSQL_DB_USER     = "MYSQL_DB_USER"
	MYSQL_DB_PASSWORD = "MYSQL_DB_PASSWORD"
)

func NewMySQLGORMConnection(ctx context.Context) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv(MYSQL_DB_USER), os.Getenv(MYSQL_DB_PASSWORD), os.Getenv(MYSQL_DB_URL),
		os.Getenv(MYSQL_DB_PORT), os.Getenv(MYSQL_DB_SCHEMA))

	logger.Info(connectionString)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db.AutoMigrate(&entity.SiteEntity{})

	if err != nil {
		return nil, err
	}
	return db, nil

}
