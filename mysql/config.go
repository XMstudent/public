package mysql

type Config interface {
	GetMysqlConfig() *ConfigMysql
}

type ConfigMysql struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	Charset      string
	MaxOpenConns int
	MaxIdleConns int
}
