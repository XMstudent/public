package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
var Service = new(Pool)

type Pool struct {
	DB *redis.Pool
}

func (rp *Pool)InitRedisPool (config *ConfigRedis)(err error){

	rp.DB = &redis.Pool{
		MaxIdle:   config.MaxIdle, /*最大的空闲连接数*/
		MaxActive: config.MaxActive, /*最大的激活连接数*/
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", fmt.Sprintf(config.Host,":",config.Port), redis.DialPassword(config.Password))
			if err != nil {
				fmt.Println("Redis错误："+err.Error())
				return nil, err
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