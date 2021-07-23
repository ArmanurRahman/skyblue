package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ArmanurRahman/skyblue/internal/config"
	"github.com/ArmanurRahman/skyblue/internal/drivers"
	"github.com/joho/godotenv"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load env file")
	}

	initiateLog()

	//initiate DB
	db, err := initiateDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.SQL.Close()

	startServer()

}

func initiateLog() {
	//initiate Log
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

}

func initiateDatabase() (*drivers.DB, error) {
	db, err := drivers.ConnectSQL(fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD")))

	if err != nil {
		log.Fatal("can't connect to database. Dying...")
		return nil, err
	}
	log.Println("connected to database")
	return db, nil
}

func startServer() {

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Start listining to port", os.Getenv("APP_PORT"))
}
