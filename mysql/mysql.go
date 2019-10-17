package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Service = new(Pool)

type Pool struct {
	DB *xorm.Engine
}

func (M *Pool) InitMysqlPool(config *ConfigMysql) (err error) {
	host := fmt.Sprintf("tcp(",config.Host,":",config.Port,")")
	maxOpenConns := config.MaxOpenConns
	maxIdleConns := config.MaxIdleConns
	dataSourceName:=fmt.Sprintf(config.Username,":",config.Password,"@",host,"/",config.Database,"?charset=",config.Charset)
	//dataSourceName := user + ":" + password + "@" + host + "/" + database + "?charset=" + charset
	M.DB, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		return err
	}
	M.DB.SetMaxOpenConns(maxOpenConns)
	M.DB.SetMaxIdleConns(maxIdleConns)

	err = M.DB.Ping()
	if err != nil {
		return err
	}
	return
}

func (M *Pool) GetClient() *xorm.Engine {
	return M.DB
}
