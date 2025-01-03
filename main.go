package main

import (
	"fmt"
	geegin "gee/geeGin"
	"html/template"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := geegin.New()
	r.Use(geegin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *geegin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	fmt.Println("start server at localhost:9999")
	r.Run(":9999")
}
