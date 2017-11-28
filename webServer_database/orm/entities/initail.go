package entities

import (
  _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
    // "github.com/go-xorm/core"
)

var engine *xorm.Engine

func Initial() {
  var err error
    engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
    // tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "orm_")
    // engine.SetTableMapper(tbMapper)
    if err != nil {
  		panic(err)
  	}
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
