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

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqlProperties.GetUser(), mysqlProperties.GetPassword(), mysqlProperties.GetUrl(),
		mysqlProperties.GetPort(), mysqlProperties.GetDBName())
	connectionLOGString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqlProperties.GetUser(), "xxxx", mysqlProperties.GetUrl(),
		mysqlProperties.GetPort(), mysqlProperties.GetDBName())

	logger.Info(connectionLOGString)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db.AutoMigrate(&entity.SiteEntity{})

	if err != nil {
		return nil, err
	}
	return db, nil

}
