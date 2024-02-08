package graphs

import "fmt"

// 1 -- 2
// |    |
// 3 -- 4

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
}

func (g *Graph) PrintGraph() {
	for k, v := range g.adjacencyList {
		fmt.Printf("%d -> %v\n", k, v)
	}
}

func (g *Graph) Bfs(startVertex int) {
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

func (g *Graph) Dfs(startVertex int) {
	visited := make(map[int]bool)
	g.DfsRecursive(startVertex, visited)
}

func (g *Graph) DfsRecursive(vertex int, visited map[int]bool) {
	if !visited[vertex] {
		fmt.Print(vertex, " ")
		visited[vertex] = true

		for _, neighbor := range g.adjacencyList[vertex] {
			if !visited[neighbor] {
				g.DfsRecursive(neighbor, visited)
			}
		}
	}
}

// func main() {
// 	g := newGraph()
// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 4)
// 	g.AddEdge(3, 4)

// 	g.PrintGraph()
// 	fmt.Println("==========DFS=========")
// 	// g.bfs(1)
// 	g.Dfs(1)
// }
