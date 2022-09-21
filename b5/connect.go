package b5

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/golangDemo5?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return engine
}
