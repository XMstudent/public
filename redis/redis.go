package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	errors "smart4s.com/public/perror"
)
var Service = new(Pool)

type Pool struct {
	DB *redis.Pool
}

func (rp *Pool)InitRedisPool (config *ConfigRedis)(err error){

	rp.DB = &redis.Pool{
		MaxIdle:   config.MaxIdle,
		MaxActive: config.MaxActive,
		Dial: func() (c redis.Conn, err error) {

			c, err = redis.Dial("tcp", fmt.Sprintf("%s:%s",config.Host,config.Port), redis.DialPassword(config.Password))
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s:%s","redis connection error:",err.Error()))
			}
			return c, nil
		},
	}
	return
}

func (rp *Pool)GetClient()redis.Conn{
	return rp.DB.Get()
}

func (rp *Pool)CloseClient(c redis.Conn)(err error){
	err = c.Close()
	if err !=nil{
		return err
	}
	return
}