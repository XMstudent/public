package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var MysqlService = new(MysqlPool)

type MysqlPool struct {
	DBMysql *xorm.Engine
}

func (M *MysqlPool) InitMysqlPool(config map[string]map[string]string) (err error) {
	host := "tcp(" + config["mysql"]["host"] + ":" + config["mysql"]["port"] + ")"
	database := config["mysql"]["database"]
	user := config["mysql"]["username"]
	password := config["mysql"]["password"]
	charset := "utf8"
	maxOpenConns := 5
	maxIdleConns := 2
	dataSourceName := user + ":" + password + "@" + host + "/" + database + "?charset=" + charset
	M.DBMysql, err = xorm.NewEngine("mysql", dataSourceName)
	fmt.Println("初始化mysql引擎")
	if err != nil {
		return err
	}
	M.DBMysql.SetMaxOpenConns(maxOpenConns)
	M.DBMysql.SetMaxIdleConns(maxIdleConns)

	err = M.DBMysql.Ping()
	if err != nil {
		return err
	}
	return
}

func (M *MysqlPool) GetClient() *xorm.Engine {
	return M.DBMysql
}
