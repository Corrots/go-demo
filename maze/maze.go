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
	if _, err := fmt.Fscanf(file, "%d %d", &row, &col); err != nil {
		panic(err)
	}
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			if _, err := fmt.Fscanf(file, "%d", &maze[i][j]); err != nil {
				panic(err)
			}
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 检查p是否越界，走出grid之外(row,上下方向)
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// 检查p是否越界，走出grid之外(col,左右方向)
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 初始化队列，并加入start
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			// maze at next is 0
			// steps at next is 0
			// next != start
			// 下一步是否越界且是不是1(墙)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// 下一步是否越界且之前以及走过了(steps对应的val不为0)
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 下一步是否是起点
			if next == start {
				continue
			}
			// 获取当前steps，并+1计入steps中next point
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			// 将下一步加入Queue
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("go-demo/maze/maze.in")
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
