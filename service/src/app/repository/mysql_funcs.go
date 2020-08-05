package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func GetConnectDataBase() *sql.DB {
	db, err := sql.Open("mysql", "root:root_password@/productdb")

	if err != nil {
		fmt.Println("panic:error")
		panic(err)
	}

	logGetConnectDataBase(err)
	return db
}
func logGetConnectDataBase(err error) {
	msg := "Welcome to service important people."
	if err != nil {
		msg = "Problems with connecting the database."
	}

	log.WithFields(log.Fields{
		"massege": msg,
		"method":  "GetConnectDataBase()",
	}).Info("Start work with db.")
}
