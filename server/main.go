package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "hola mundo")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
