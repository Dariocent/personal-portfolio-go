package main

import (
    "github.com/Dariocent/personal-portfolio-go/blog"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
		return c.File("templates/index.html")
	})

    e.GET("/base", func(c echo.Context) error {
		return c.File("templates/base.html")
    })

    e.GET("/blog", func(c echo.Context) error {
		return c.File("templates/blog.html")
	})

    e.Static("/static", "static")

    api := e.Group("/api")
    blog.SetupRoutes(api.Group("/blog"))
    
    e.Logger.Fatal(e.Start(":8080"))

    e.Start(":8080")
}