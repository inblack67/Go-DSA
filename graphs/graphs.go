package graphs

import "fmt"

// 1 -- 2
// |    |
// 3 -- 4

type Graph struct {
	adjacencyList map[int][]int
}

func newGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (g *Graph) addEdge(v1, v2 int) {
	g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
}

func (g *Graph) printGraph() {
	for k, v := range g.adjacencyList {
		fmt.Printf("%d -> %v\n", k, v)
	}
}

func main() {
	g := newGraph()
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 4)
	g.addEdge(3, 4)

	g.printGraph()
}
