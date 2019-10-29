package redis

import (
	"fmt"
	errors "github.com/XMstudent/public/perror"
	"github.com/gomodule/redigo/redis"
)

var Service = new(Pool)

type Pool struct {
	DB *redis.Pool
}

func (rp *Pool) InitRedisPool(config Config) (err error) {
	redisConfig:=config.GetRedisConfig()
	rp.DB = &redis.Pool{
		MaxIdle:   redisConfig.MaxIdle,
		MaxActive: redisConfig.MaxActive,
		Dial: func() (c redis.Conn, err error) {

			c, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port), redis.DialPassword(redisConfig.Password))
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s:%s", "redis connection error:", err.Error()))
			}
			return c, nil
		},
	}
	return
}

func (rp *Pool) GetClient() redis.Conn {
	return rp.DB.Get()
}

func (rp *Pool) CloseClient(c redis.Conn) (err error) {
	err = c.Close()
	if err != nil {
		return err
	}
	return
}
