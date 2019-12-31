package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Fscanf(file, "%d") //抛弃换行符
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, "%d") // 抛弃换行符
	}

	return maze
}

type point struct {
	i, j int
}

var search = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(maze [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(maze) || p.j < 0 || p.j >= len(maze[p.i]) {
		return 0, false
	}
	return maze[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for row := range steps {
		steps[row] = make([]int, len(maze[row]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		curPoint := Q[0]
		Q = Q[1:]

		if curPoint == end {
			break
		}

		//探索当前节点的上下左右节点
		for _, direction := range search {
			next := curPoint.add(direction)

			//判断next是否可以走

			//节点在maze里面，且可以走
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			//节点没有走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			//回到起始节点
			if next == start {
				continue
			}

			val, _ = curPoint.at(steps)
			steps[next.i][next.j] = val + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("src/maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	steps := walk(maze, point{0, 0}, point{2, 2})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
