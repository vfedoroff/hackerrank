package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node is a graph node
type Node struct {
	ID    int
	Links []*Node `json:"-"`
}

func (n *Node) String() string {
	js, _ := json.Marshal(n)
	return fmt.Sprintf("%s", js)
}

// Graph is a graph
type Graph struct {
	nodes map[int]*Node
}

func (g *Graph) addNode(id int) {
	_, srcOk := g.nodes[id]
	if !srcOk {
		g.nodes[id] = &Node{ID: id, Links: []*Node{}}
	}
}

func (g *Graph) addEdge(src, dst int) {
	srcNode, srcOk := g.nodes[src]
	dstNode, dstOk := g.nodes[dst]
	if !srcOk {
		g.nodes[src] = &Node{ID: src, Links: []*Node{}}
	}
	if !dstOk {
		g.nodes[dst] = &Node{ID: dst, Links: []*Node{}}
	}
	srcNode.Links = append(srcNode.Links, dstNode)
	g.nodes[src] = srcNode
}

func (g *Graph) Bfs(start int) map[int]int {
	// Mark the current node as visited and enqueue it
	visited := make(map[int]bool)
	distance := make(map[int]int)
	for x := 1; x <= len(g.nodes); x++ {
		distance[x] = -1
	}
	visited[start] = true
	distance[start] = 0
	queue := []int{start}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range g.nodes[u].Links {
			if !visited[v.ID] {
				visited[v.ID] = true
				queue = append(queue, v.ID)
				distance[v.ID] = distance[u] + 6
			}
		}
	}
	return distance
}

func newGraph() *Graph {
	return &Graph{nodes: map[int]*Node{}}
}

func (g *Graph) String() {
	s := ""
	for _, node := range g.nodes {
		s += node.String() + " -> "
		for _, near := range node.Links {
			s += near.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	q, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}
	for i := 0; i < q; i++ {
		graph := newGraph()
		scanner.Scan()
		str := scanner.Text()
		var n, m int
		fmt.Sscanf(str, "%d %d", &n, &m)
		for x := 1; x <= n; x++ {
			graph.addNode(x)
		}
		for l := 0; l < m; l++ {
			var u, v int
			scanner.Scan()
			str := scanner.Text()
			fmt.Sscanf(str, "%d %d", &u, &v)
			// add each edge to the graph
			graph.addEdge(u, v)
			graph.addEdge(v, u)
		}
		scanner.Scan()
		str = scanner.Text()
		s, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		dist := graph.Bfs(s)
		out := []int{}
		for x := 1; x <= n; x++ {
			if x != s {
				d, ok := dist[x]
				if ok {
					out = append(out, d)
					continue
				}
				out = append(out, -1)
			}
		}
		fmt.Println(strings.Trim(fmt.Sprint(out), "[]"))
	}
}
