// db.go
package main

import (
	"database/sql"
	"fmt"
	"os"

	mysqldriver "github.com/go-sql-driver/mysql"
)

var (
	CONNECTION   *sql.DB
	 
)
 
func init() {
    
	mysqlConfig:= &mysqldriver.Config{
     
		User:                  "root",
		Passwd:                "MaMa1ghada",
		Net:                   "tcp",
		Addr:                  os.Getenv("MYSQL_HOST")+ ":"+os.Getenv("MYSQL_PORT"),
		AllowNativePasswords: false,
 		ParseTime:            true,
	}
	

	connection, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		fmt.Println(err,2222222)
	}

	_, err = connection.Exec(`CREATE DATABASE IF NOT EXISTS internship`)
	if err != nil {
		fmt.Println(err,1111111111)
	}

	mysqlConfig.DBName = `internship`

	if err = connection.Close(); err != nil {
		panic(err)
	}

	connection, err = sql.Open("mysql",mysqlConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	CONNECTION  = connection

	schema()
}

 

func schema() {
	_, err := CONNECTION.Exec(`CREATE TABLE IF NOT EXISTS stuff (
		id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		created_at DATETIME NOT NULL
	)`)

	if err != nil {
		panic(err)
	}
}
