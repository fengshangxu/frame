package main

import (
	"net/http"
)

import (
	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1> 123 </h1>")
	})

	r.GET("/string", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/json", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.Query("username"),
			"password": c.Query("password"),
		})
	})

	r.Run(":9999")
}