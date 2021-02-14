package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer a.DB.Close()

	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connection. Successfully!")

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
