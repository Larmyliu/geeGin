package geegin

import (
	"net/http"
)

type HandlerFunc func(*Context)
type Engine struct {
	*RouterGroup
	Router *router
	Groups []*RouterGroup
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{Router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
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
