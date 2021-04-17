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

	//instance echo instance and dbs
	e := echo.New()
	store := cache.SessionsClient()

	// middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
		AllowCredentials: true,
	}))
	e.Pre(cache.Sessions("qid", *store))
	e.Use(middleware.RemoveTrailingSlash())

	// set routes
	routes.Router(e)

	// start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
