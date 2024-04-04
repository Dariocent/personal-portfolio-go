package blog

type Article struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

// Sample data for testing
var Articles = []Article{
    {ID: 1, Title: "First Post", Content: "Hello, this is my first post."},
    {ID: 2, Title: "Second Post", Content: "Another day, another post."},
}
