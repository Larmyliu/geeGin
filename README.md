# geeGin

这是一个一遍阅读Gin源码，一遍动手实践的项目

主要参考https://geektutu.com/post/gee.html

也会分享一些开发过程中，结合源码阅读后的一些理解

## Day1

主要实现了，Engine、route、Context的拆分

Engine依赖于原生"net/http"包，通过实现Handler的接口，启动http服务，搭配route实现挂载路由

```golang
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

route主要就是实现路由注册，method+path确定唯一的key值，第一天Route是用MAP来实现的，实际上GIN是用树实现的，保存树用的是切片  type methodTrees []methodTree

Context就是定义上下文，和Route拆分开来，对ResponseWriter和Request进行封装，并在里面实现一些简易方法，比如Query,SetStatus等，Context是跟随一个请求完整的生命周期的，后续补充c.Get(),c.Set()


