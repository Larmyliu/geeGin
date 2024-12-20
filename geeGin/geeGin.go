package geegin

import (
	"net/http"
)

type HandlerFunc func(*Context)
type Engine struct {
	Router *router //这里Router是用MAP来实现的，实际上GIN是用树实现的，保存树用的是切片  type methodTrees []methodTree
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{Router: newRouter()}
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.Router.handle(c)
}

// 添加路由，这里只支持了GET和POST，实际上GIN支持9种请求方式：GET、POST、PUT、DELETE、OPTIONS、HEAD、TRACE、PATCH、CONNECT
// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.Router.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.Router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
