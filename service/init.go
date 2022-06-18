package service

import (
	"blog-api/model"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	mysqlRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	// docker-composeで動かす場合、hostはlocalhostではなくcontainer_nameになる
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	DsName := "root:" + mysqlRootPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/blog?charset=utf8"
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
