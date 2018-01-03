package entities

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // .
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("mysql", "root:ren avril@tcp(127.0.0.1:3306)/xorm_UserInfo?charset=utf8&parseTime=true")
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{}, "xorm_")
	engine.SetTableMapper(tableMapper)

	err := engine.Sync(new(UserInfo))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
