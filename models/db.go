package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

func Database() {
	_ = godotenv.Load("db.env")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")
	port := os.Getenv("port")

	// dburl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=% sslmode=disable", host, port, user, password, dbname)
	dburl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)


	Db, err = sql.Open("postgres", dburl)
	if err != nil {
		fmt.Println("error opening db")
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	// defer Db.Close()

	fmt.Println("Connected successfully")
}
