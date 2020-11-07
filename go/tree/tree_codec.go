package main

import (
	"container/list"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	q := list.New()
	q.PushFront(root)
	result := []string{c.toStr(root)}
	for q.Len() > 0 {
		node := q.Remove(q.Back()).(*TreeNode)
		result = append(result, c.toStr(node.Left), c.toStr(node.Right))
		if node.Left != nil {
			q.PushFront(node.Left)
		}
		if node.Right != nil {
			q.PushFront(node.Right)
		}
	}
	i := len(result) - 1
	for ; i > 0 && result[i] == "null"; i-- {
	}
	res := "[" + strings.Join(result[:i+1], ",") + "]"
	return res
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	N := len(data)
	if N < 3 || data[0] != '[' || data[N-1] != ']' {
		return nil
	}
	vals := strings.Split(data[1:N-1], ",")
	if len(vals) == 0 {
		return nil
	}
	root, i := c.toNode(vals[0]), 1
	q := list.New()
	q.PushFront(root)
	for q.Len() > 0 && i < len(vals) {
		node := q.Remove(q.Back()).(*TreeNode)
		node.Left = c.toNode(vals[i])
		if i+1 < len(vals) {
			node.Right = c.toNode(vals[i+1])
		}
		i += 2
		if node.Left != nil {
			q.PushFront(node.Left)
		}
		if node.Right != nil {
			q.PushFront(node.Right)
		}
	}
	return root
}

func (c *Codec) toStr(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return strconv.Itoa(node.Val)
}
func (c *Codec) toNode(s string) *TreeNode {
	if s == "null" {
		return nil
	}
	v, _ := strconv.Atoi(s)
	return &TreeNode{v, nil, nil}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
