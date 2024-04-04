package blog

import (
    "html/template"
    "github.com/labstack/echo/v4"
)

func GetArticles(c echo.Context) error {
    tmpl := `
    <ul id="articles-list">
        {{range .}}
            <div class="shadow-1-strong bg-white my-5 p-5" id="articles">
            <li>
                <h3>{{.Title}}</h3>
                <p>{{.Content}}</p>
            </li>
            </div>
        {{end}}
    </ul>
    `
    t := template.Must(template.New("articles").Parse(tmpl))
    
    return t.Execute(c.Response().Writer, Articles)
}

