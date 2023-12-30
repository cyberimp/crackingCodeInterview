package main

import (
	"encoding/json"
)

func isPassable(edges [][]int, s int, t int) bool {
	visited := make(map[int]bool)
	graph := make(map[int][]int)

	var start, stop int
	for _, e := range edges {
		start = e[0]
		stop = e[1]
		graph[start] = append(graph[start], stop)
	}

	queue := []int{s}

	var cur int

	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		if !visited[cur] {
			visited[cur] = true
			queue = append(queue, graph[cur]...)
		}
	}
	return visited[t]
}

type Node struct {
	Left  *Node `json:"left,omitempty"`
	Right *Node `json:"right,omitempty"`
	Value int   `json:"value"`
}

func makeNode(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	if len(values) == 1 {
		return &Node{nil, nil, values[0]}
	}
	middle := len(values) / 2
	return &Node{makeNode(values[:middle]), makeNode(values[middle+1:]), values[middle]}

}

func makeBST(values []int) string {
	root := makeNode(values)
	result, _ := json.Marshal(root)
	return string(result)
}

func binaryTreeIntoLists(tree string) [][]int {
	var root Node
	_ = json.Unmarshal([]byte(tree), &root)
	result := make([][]int, 0)
	var curRow []*Node
	var nextRow = []*Node{&root}
	deep := 0
	for len(nextRow) > 0 {
		curRow = nextRow
		nextRow = make([]*Node, 0)
		result = append(result, make([]int, 0))

		for _, node := range curRow {
			result[deep] = append(result[deep], node.Value)
			if node.Left != nil {
				nextRow = append(nextRow, node.Left)
			}
			if node.Right != nil {
				nextRow = append(nextRow, node.Right)
			}
		}

		deep++
	}
	return result
}
