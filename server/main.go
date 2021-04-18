package main

import (
	"CodePonder/cache"
	"CodePonder/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to read env")
	}

	//instance echo instance and redis client
	e := echo.New()
	store := cache.SessionsClient()

	// middlewares
	e.Use(routes.CORSConfig())
	e.Use(cache.Sessions("qid", *store))
	e.Use(routes.IsAuth)
	e.Use(middleware.RemoveTrailingSlash())

	// set routes
	routes.Router(e)

	// start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
