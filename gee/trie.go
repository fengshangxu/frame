package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

//匹配节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	nodes := n.matchChildren(part)
	for _, node := range nodes {
		result := node.search(parts, height+1)
		if nil != result {
			return result
		}
	}
	return nil
}

//匹配节点，用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
