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

## Day2

主要实现前缀树，前缀树通过给每个不同的method都对应一个树，根节点为共有部分，来进行路径匹配

其中最有意思的就是参数动态匹配，包括参数匹配（:）和通配符匹配（*），这也是项目的核心代码

匹配思路如下

1. 路径解析：先传入完整的路径`pattern`, 会把完整的路径根据`/`分割，分割出来的每个部分开始递归创建前缀树的Node节点，例如，路径 /p/:lang/doc 会被解析为 ["p", ":lang", "doc"]。
2. 节点插入: 解析得到的路径开始**递归**插入进前缀树中，如果路径没有在前缀树中出现过，就创建一个节点，如果匹配第一个字符是'*'或者':', 代表参数匹配，isWild为true。递归结束条件就是len(parts) == height，这时候把完整路径给节点赋值一下
3. 路径匹配：将请求路径再次解析成部分（根据/分割），通过search方法沿着前缀树节点进行匹配。如果发现动态参数（包含*或者/），把这些特殊参数去掉，并存到param中，后面可以在c.Param中查询

跟gin框架比，这里的实现是比较糙的，gin的源代码对前缀树还有功能和算法上的优化等等，但是通过这个实践也是可以很好的了解到前缀树的思想，然后去读源码的时候才能更理解框架的实现思路

## Day2 -1
给路由建立分组，创建一个RouterGroup类，里面有路径前缀，中间件回调（后续实现）。

这里有个很有意思的事情就是，评论区针对RouterGroup 是否应该嵌套 Engine有激烈的讨论，我这边认为是需要的，翻看gin的源码也是有引用的
```go
// RouterGroup is used internally to configure router, a RouterGroup is associated with
// a prefix and an array of handlers (middleware).
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}
```
这里表达一下自己粗浅的看法，如果有其他想法也欢迎一起讨论。

我认为，组这个概念是一个语法糖，可以方便我们对具有相同前缀的api做一些特性处理，比如挂载统一的鉴权中间件或者日志中间件。如果没有这个组的概念，Engine一样也是可以实现的，在`*router`上继续扩middlewares就行了。因为添加路由最底层的实现方法还是`*router.addRoute`。

现在是为了使类抽象的更加好，所以抽出了一层`RouterGroup`，实际调用的底层还是`*router.addRoute`，作为底层实现是典型的面向接口编程思想。这样即使底层的路由实现发生改变，RouterGroup的接口也可以保持稳定，对外提供统一接口。