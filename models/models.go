package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:bxsec@tcp(127.0.0.1:33060)/godSkills?charset=utf8")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//f, err := os.Create("godSkills_sql.log")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//engine.SetLogger(xorm.NewSimpleLogger(f))

	engine.SetMaxOpenConns(10)
}

func Engine() *xorm.Engine {
	return engine
}
