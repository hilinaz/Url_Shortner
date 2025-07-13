package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
var DB *sql.DB
func Connect(){
	// load the env
	err := godotenv.Load()
	if err!=nil{
		log.Fatal(".env file not found")

	}
	// read env
	user:= os.Getenv("DB_USER")
	pass:= os.Getenv("DB_PASS")
	host:= os.Getenv("DB_HOST")
	port:= os.Getenv("DB_PORT")
	dbName:= os.Getenv("DB_NAME")

	dsnNoDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)

	DB,err=sql.Open("mysql",dsnNoDB)
	if err!=nil{
		log.Fatal("Database connection failed")
	}


	// create db
	_,err = DB.Exec("CREATE DATABASE IF NOT EXISTS " +dbName)
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}
	log.Println("Database ensured:", dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the actual database:", err)
	}
	err=DB.Ping()
	if err!=nil{
		log.Fatal("connection failed")
	}

	log.Print("Successfully connected to the databse")
	

}
