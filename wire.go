//+build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"

	"github.com/Alkemic/gowroc-di-project/blog"
	"github.com/Alkemic/gowroc-di-project/handler"
	"github.com/Alkemic/gowroc-di-project/repository"
)

func InitializeHttpHandler() handler.HttpHandler {
	wire.Build(handler.NewHandler, blog.NewBlogService, repository.NewRepository, openDB, getDBDSN)
	return &http.ServeMux{}
}
