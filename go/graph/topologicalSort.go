package main

import "fmt"

type Node struct {
	val      string
	children []*Node
}

func (n *Node) add(child *Node) *Node {
	n.children = append(n.children, child)
	return n
}
func (n *Node) addStr(c string) *Node {
	child := newNode(c)
	n.add(child)
	return child
}
func (n *Node) get(s string) *Node {
	for _, c := range n.children {
		if c.val == s {
			return c
		}
	}
	return nil
}

func newNode(s string) *Node {
	return &Node{val: s}
}

func tpSort(root *Node) string {
	tmp := dfs(root, make(map[*Node]bool))
	n := len(tmp)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = tmp[n-i-1]
	}
	return string(res)
}
func dfs(root *Node, flag map[*Node]bool) string {
	if root == nil || flag[root] {
		return ""
	}
	flag[root] = true
	result := ""
	for _, child := range root.children {
		if child != nil && !flag[child] {
			result += dfs(child, flag)
		}
	}
	return result + root.val
}

func main() {
	// A +-> B -+> D -+> E --> F --> G
	//   +-> C -+     |
	//   +------------+
	root := newNode("A")
	root.addStr("B").addStr("D").
		addStr("E").addStr("F").addStr("G")
	root.addStr("C")
	d := root.get("B").get("D")
	e := d.get("E")
	root.add(e)
	fmt.Println(tpSort(root))
}
