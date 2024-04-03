package main

import (

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
		return c.File("templates/index.html")
	})

	e.Static("/static", "static")

    e.Start(":8080")
}