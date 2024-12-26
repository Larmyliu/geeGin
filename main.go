package main

import (
	"fmt"
	geegin "gee/geeGin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() geegin.HandlerFunc {
	return func(c *geegin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.HttpStatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	engine := geegin.New()
	engine.Use(geegin.Logger())
	engine.GET("/", func(c *geegin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	v1 := engine.Group("/v1")
	v1.GET("/hello", func(c *geegin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	v1.GET("/hello/:name", func(c *geegin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v2 := engine.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/assets/*filepath", func(c *geegin.Context) {
		c.JSON(http.StatusOK, geegin.H{"filepath": c.Param("filepath")})
	})
	err := engine.Run(":9999")
	if err != nil {
		panic(err)
	}
	fmt.Println("server started at localhost:9999")
}
