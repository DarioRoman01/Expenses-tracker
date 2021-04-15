package main

import (
	"CodePonder/cache"
	"CodePonder/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//instance echo instance and dbs
	e := echo.New()
	store := cache.SessionsClient()

	// middlewares
	e.Use(cache.Sessions("qid", *store))
	e.Use(middleware.RemoveTrailingSlash())

	routes.Router(e)
	e.Logger.Fatal(e.Start(":1323"))
}
