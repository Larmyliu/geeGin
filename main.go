package main

import (
	"fmt"
	geegin "gee/geeGin"
	"net/http"
)

func main() {
	engine := geegin.New()
	engine.GET("/", func(c *geegin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	engine.GET("/hello", func(c *geegin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	engine.GET("/hello/:name", func(c *geegin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	engine.GET("/assets/*filepath", func(c *geegin.Context) {
		c.JSON(http.StatusOK, geegin.H{"filepath": c.Param("filepath")})
	})
	err := engine.Run(":9999")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started at localhost:9999")
}
