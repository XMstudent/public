package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var Service = new(Pool)

type Pool struct {
	DB *xorm.Engine
}

func (M *Pool) InitMysqlPool(config Config) (err error) {
	mysqlConfig:=config.GetMysqlConfig()
	host := fmt.Sprintf("tcp(%s:%s)",mysqlConfig.Host,mysqlConfig.Port)
	maxOpenConns := mysqlConfig.MaxOpenConns
	maxIdleConns := mysqlConfig.MaxIdleConns
	dataSourceName :=fmt.Sprintf("%s:%s@%s/%s?charset=%s",mysqlConfig.Username,mysqlConfig.Password,host,mysqlConfig.Database,mysqlConfig.Charset)
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
