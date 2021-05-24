package dbUtil

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:bxsec@/godSkills?charset=utf8")
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