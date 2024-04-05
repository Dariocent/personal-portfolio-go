package blog

import (
    "html/template"
    "github.com/labstack/echo/v4"
    "strconv"
    "net/http"
)

func GetArticles(c echo.Context) error {
    t, err := template.ParseFiles("templates/articles.html")
    if err != nil {
        return err
    }
    
    return t.Execute(c.Response().Writer, Articles)
}

func GetArticle(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return err
    }
    article := Articles[id]
    t, err := template.ParseFiles("templates/article.html")
    if err != nil {
        return err
    }
    
    return t.Execute(c.Response().Writer, article)
}

func DeleteArticle(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return err
    }
    Articles = append(Articles[:id], Articles[id+1:]...) //append everything before and after the article to be deleted
    return c.JSON(http.StatusOK, Articles)
}

func CreateArticle(c echo.Context) error {
    article := new(Article)
    article.ID = 3
    article.Title = "Third Post"
    article.Content = "Yet another post."
    if err := c.Bind(article); err != nil {
        return err
    }
    Articles = append(Articles, *article)
    return c.JSON(http.StatusCreated, article)
}
