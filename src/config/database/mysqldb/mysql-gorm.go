package mysqldb

import (
	"backend-site/src/config/aws.secrets/model"
	"backend-site/src/config/logger"
	"backend-site/src/model/repository/entity"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLGORMConnection(ctx context.Context, mysqlProperties model.MySQLPropertiesInterface) (*gorm.DB, error) {

	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	//	os.Getenv(MYSQL_DB_USER), os.Getenv(MYSQL_DB_PASSWORD), os.Getenv(MYSQL_DB_URL),
	//	os.Getenv(MYSQL_DB_PORT), os.Getenv(MYSQL_DB_SCHEMA))
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysqlProperties.GetUrl(), mysqlProperties.GetPassword(), mysqlProperties.GetUrl(),
		mysqlProperties.GetPort(), mysqlProperties.GetDBName())

	logger.Info(connectionString)

	logger.Info(connectionString)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db.AutoMigrate(&entity.SiteEntity{})

	if err != nil {
		return nil, err
	}
	return db, nil

}
