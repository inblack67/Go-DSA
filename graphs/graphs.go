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

func (g *Graph) bfs(startVertex int) {
	visited := make(map[int]bool)
	queue := []int{startVertex}
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		if !visited[vertex] {
			fmt.Print(vertex, " ")
			visited[vertex] = true
			for _, neighbor := range g.adjacencyList[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

}

func (g *Graph) dfs(startVertex int) {
	visited := make(map[int]bool)
	g.dfsRecursive(startVertex, visited)
}

func (g *Graph) dfsRecursive(vertex int, visited map[int]bool) {
	if !visited[vertex] {
		fmt.Print(vertex, " ")
		visited[vertex] = true

		for _, neighbor := range g.adjacencyList[vertex] {
			if !visited[neighbor] {
				g.dfsRecursive(neighbor, visited)
			}
		}
	}
}

func main() {
	g := newGraph()
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 4)
	g.addEdge(3, 4)

	g.printGraph()
	fmt.Println("==========DFS=========")
	// g.bfs(1)
	g.dfs(1)
}
