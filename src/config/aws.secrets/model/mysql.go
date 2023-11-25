package model

type mysqlProperties struct {
	url      string
	port     uint64
	dbname   string
	user     string
	password string
}

func (m *mysqlProperties) GetUrl() string {
	return m.url
}

func (m *mysqlProperties) GetPort() uint64 {
	return m.port
}

func (m *mysqlProperties) GetDBName() string {
	return m.dbname
}

func (m *mysqlProperties) GetUser() string {
	return m.user
}

func (m *mysqlProperties) GetPassword() string {
	return m.password
}
