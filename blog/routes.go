package blog

import (
    "github.com/labstack/echo/v4"
)

func SetupRoutes(g *echo.Group) {
    g.GET("/articles", GetArticles)
    g.GET("/articles/:id", GetArticle)
    g.DELETE("/articles/:id", DeleteArticle)
    g.POST("/articles", CreateArticle)
}