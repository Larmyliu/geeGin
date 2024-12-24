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

