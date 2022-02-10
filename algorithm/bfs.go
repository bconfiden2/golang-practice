package main

import "fmt"

func bfs(start, end int, graph [][]int) (int) {
	q := make([]int, 1)
	q[0] = start
	visited := make([]int, len(graph))
	visited[start] = 1

	for len(q) > 0 {
		cur := q[0]
		if cur == end {
			return visited[cur]
		}
		q = q[1:]
		for i := 0 ; i < len(graph[cur]) ; i++ {
			nxt := graph[cur][i]
			if visited[nxt] == 0 {
				visited[nxt] = visited[cur] + 1
				q = append(q, nxt)
			}
		}
	}
	return 0
}

func main() {
	graph := make([][]int, 4)
	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)
	graph[1] = append(graph[1], 3)
	
	fmt.Println(bfs(0, 3, graph))
}