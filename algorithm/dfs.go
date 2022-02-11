package main

import (
	"fmt"
)

var R, C int
var dir [4]int = [4]int{1,-1,0,0}

func main() {
	R, C = 8, 8
	graph := [][]int{
		{1,0,0,1,1,1,0,1},
		{0,0,0,0,1,1,1,0},
		{1,1,1,0,0,0,0,1},
		{0,1,0,1,1,1,1,1},
		{0,0,0,1,0,0,0,1},
		{1,0,0,1,0,1,0,1},
		{1,0,0,0,1,0,0,1},
		{1,1,1,1,0,0,0,1},
	}
	cnt := 0
	visited := make([][]bool, 8)
	for i := 0 ; i < 8 ; i++ {
		visited[i] = make([]bool, 8)
	}

	for i := 0 ; i < 8 ; i++ {
		for j := 0 ; j < 8 ; j++ {
			if visited[i][j] == false && graph[i][j] == 1 {
				dfs(graph, i, j, visited)
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func dfs(graph [][]int, r, c int, visited [][]bool) {
	visited[r][c] = true
	for i := 0 ; i < 4 ; i++ {
		nr, nc := r + dir[i], c + dir[3-i]
		if 0 <= nr && nr < R && 0 <= nc && nc < C {
			if visited[nr][nc] == false && graph[nr][nc] == 1 {
				dfs(graph, nr, nc, visited)
			}
		}
	}
}