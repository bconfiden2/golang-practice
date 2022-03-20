package main

import "fmt"

var (
	n, m int
	u, v int
	graph [][]int
)

func main() {
	n, m = 4, 3

	graph = make([][]int, n+1)
	for i := 0 ; i < n+1 ; i++ {
		graph[i] = make([]int, n+1)
	}
	for r := 1 ; r <= n ; r++ {
		for c := 1 ; c <= n ; c++ {
			graph[r][c] = 1000000
		}
	}

	graph[1][2] = 1
	graph[2][1] = 1
	graph[2][3] = 1
	graph[3][2] = 1
	graph[3][4] = 1
	graph[4][3] = 1

	for x := 1 ; x <= n ; x++ {
		graph[x][x] = 0
		for r := 1 ; r <= n ; r++ {
			for c := 1 ; c <= n ; c++ {
				if graph[r][c] > graph[r][x] + graph[x][c] {
					graph[r][c] = graph[r][x] + graph[x][c]
				}
			}
		}
	}

	for r := 1 ; r <= n ; r++ {
		for c := 1 ; c <= n ; c++ {
			if graph[r][c] == 1000000 {
				fmt.Print(0, " ")
			} else {
				fmt.Print(graph[r][c], " ")
			}
		}
		fmt.Println()
	}
}