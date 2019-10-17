package redis

type Config interface {
	GetRedisConfig() *ConfigRedis
}

type ConfigRedis struct {
	Host      string
	Port      string
	Password  string
	MaxIdle   int
	MaxActive int
}
