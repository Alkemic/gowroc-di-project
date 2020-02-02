package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Alkemic/gowroc-di-project/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	container := buildContainer()
	err := container.Invoke(func(httpHandler handler.HttpHandler) {
		srv := http.Server{Addr: ":8080", Handler: httpHandler}
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("server failed: %v", err)
		}
	})

	if err != nil {
		panic(err)
	}
}

type DBDSN string

func openDB(dsn DBDSN) *sql.DB {
	db, err := sql.Open("mysql", string(dsn))
	if err != nil {
		log.Fatalln("error connecting to database", err)
	}
	return db
}

func getDBDSN() DBDSN {
	return "root:root@/blog?parseTime=true"
}
