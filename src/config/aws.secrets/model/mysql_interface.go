package model

func NewMySQLProperties(url string, port uint64, dbname string, user string, password string) MySQLPropertiesInterface {
	return &mysqlProperties{
		url:      url,
		port:     port,
		dbname:   dbname,
		user:     user,
		password: password,
	}
}

type MySQLPropertiesInterface interface {
	GetUrl() string

	GetPort() uint64

	GetDBName() string

	GetUser() string

	GetPassword() string
}
