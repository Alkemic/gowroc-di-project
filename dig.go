package main

import (
	"go.uber.org/dig"

	"github.com/Alkemic/gowroc-di-project/blog"
	"github.com/Alkemic/gowroc-di-project/handler"
	"github.com/Alkemic/gowroc-di-project/repository"
)

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(getDBDSN)
	container.Provide(openDB)
	container.Provide(handler.NewHandler)
	container.Provide(repository.NewRepository)
	container.Provide(blog.NewBlogService)
	return container
}
