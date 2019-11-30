package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Alkemic/gowroc-di-project/blog"
	"github.com/Alkemic/gowroc-di-project/handler"
	"github.com/Alkemic/gowroc-di-project/repository"
)

func main() {
	db := openDB("root:root@/blog?parseTime=true")
	repo := repository.NewRepository(db)
	blogService := blog.NewBlogService(repo)
	httpHandler := handler.NewHandler(blogService)
	srv := http.Server{Addr: ":8080", Handler: httpHandler}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln("server failed: %v", err)
	}
}

func openDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("error connecting to database", err)
	}
	return db
}
