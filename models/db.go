package models

import (
	"database/sql"
	"fmt"
	"tiny-blog/configs"
	//log "tiny-blog/middlewares"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Db is mysql
var Db *sql.DB

func DbInit() {
	var err error
	fmt.Println(configs.Conf.Db)
	connstr := fmt.Sprintf("%s:%v@tcp(%v:%v)/%v?charset=utf8mb4",
		configs.Conf.Db.User,
		configs.Conf.Db.Password,
		configs.Conf.Db.Host,
		configs.Conf.Db.Port,
		configs.Conf.Db.Name)
	Db, err = sql.Open("mysql", connstr)
	fmt.Println(connstr)
	//log.Logger.Info(connstr)
	if err != nil {
		//log.Logger.Panicln("err:", err.Error())
		fmt.Println(err.Error())
		return
	}
	Db.SetMaxOpenConns(configs.Conf.Db.Maxconns)
	Db.SetMaxIdleConns(configs.Conf.Db.Maxconns)
}
