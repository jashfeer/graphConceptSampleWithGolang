package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 1},
	}

	fmt.Println(riverSize(matrix))

}


func riverSize(matrix [][]int) []int {
	sizes := []int{}

	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			} else {
				sizes = traversed(i, j, matrix, visited, sizes)
			}
		}
	}

	return sizes
}
func traversed(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {

	currentRiversize := 0

	nodesToExplore := [][]int{{i, j}}

	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j = currentNode[0], currentNode[1]

		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		currentRiversize++
		nearest := nearestNodes(i, j, matrix, visited)

		nodesToExplore = append(nodesToExplore, nearest...)

	}
	if currentRiversize > 0 {

		sizes = append(sizes, currentRiversize)
	}

	return sizes
}
func nearestNodes(i, j int, matrix [][]int, visited [][]bool) [][]int {

	nearest := [][]int{}

	if i > 0 && !visited[i-1][j] {

		nearest = append(nearest, []int{i - 1, j})
	}

	if i < len(matrix)-1 && !visited[i+1][j] {

		nearest = append(nearest, []int{i + 1, j})
	}
	if j > 0 && !visited[i][j-1] {

		nearest = append(nearest, []int{i, j - 1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {

		nearest = append(nearest, []int{i, j + 1})
	}

	return nearest
}

