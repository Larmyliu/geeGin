package main

import (
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

	engine.Run(":9999")
}
