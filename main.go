package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type blog struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

var blogs = []blog{
	{ID: "1", Title: "First blog", Description: "Example content", Slug: "first-blog"},
	{ID: "2", Title: "Second blog", Description: "Example content", Slug: "second-blog"},
	{ID: "3", Title: "Third Blog", Description: "Example content", Slug: "third-blog"},
}

func main() {
	router := gin.Default()
	router.GET("/blogs", getBlogs)
	router.GET("/blog/:slug", getBlogBySlug)
	router.POST("/blogs", postBlogs)

	router.Run("localhost:8080")
}

func getBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogs)
}

func postBlogs(c *gin.Context) {
	var newBlog blog

	if err := c.BindJSON(&newBlog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	for _, a := range blogs {
		if a.ID == newBlog.ID {
			c.IndentedJSON(http.StatusUnprocessableEntity, blogs)
			return
		}
	}

	blogs = append(blogs, newBlog)
	c.IndentedJSON(http.StatusCreated, newBlog)
}

func getBlogBySlug(c *gin.Context) {
	slug := c.Param("slug")

	for _, a := range blogs {
		if a.Slug == slug {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog post not found"})
}
