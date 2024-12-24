package geegin

import "strings"

type Node struct {
	Pattern  string  // 待匹配路由
	Part     string  // 路由中的一部分
	Children []*Node // 子路径
	IsWild   bool    // 是否通配符节点, part 含有 : 或 * 时为true
}

func (n *Node) matchChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

func (n *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// 没匹配到就创建一个
		child = &Node{
			Part:   part,
			IsWild: part[0] == ':' || part[0] == '*',
		}
		n.Children = append(n.Children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *Node) search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
