package service

import (
	"blog-api/model"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	// DsName := "root:mysql_pass2021@tcp(localhost:3306)/blog?charset=utf8"
	DsName := "root:password@tcp(blog-db:3306)/blog?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.SetMaxOpenConns(2)
	DbEngine.ShowSQL(true)
	err = DbEngine.Sync2(new(model.User))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("init data base ok")
}
