package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day12() {
	inputText, err := os.Open("input/day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	heightMap := make([][]rune, 0)
	visited := make(map[string]int)
	rowCount := 0
	startX := 0
	startY := 0

	//nodes := make(map[Position]mapNode) //le guin i thought about making an slice of mapNodes and with "location Position" but i wasnt sure how to traverse it

	scanner := bufio.NewScanner(inputText)

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "S") {
			startX = strings.Index(scanner.Text(), "S")
			startY = rowCount
		}
		row := []rune(scanner.Text())
		heightMap = append(heightMap, row)
		rowCount++
	}

	for index := range heightMap {
		fmt.Println(string(heightMap[index]))
	}

	//generate possible path coords
	//if they contain -1 or parent skip them
	//iterate to that path
	//fmt.Println(startX, startY)
	//fmt.Println(len(heightMap))

	fmt.Println(string(heightMap[1][3]))
	start := mapNode{nil, startX, startY, 'a'}
	fmt.Println(start.checkPath(heightMap, visited, 0))

}

type mapNode struct {
	parent *mapNode

	x     int
	y     int
	value rune
}

func inBounds(x int, y int, width int, height int) bool {
	if x >= 0 && x < width && y >= 0 && y < height {
		return true
	}
	return false
}

func notTooHigh(current rune, target rune) bool {
	if current == 'z' && target == 'E' {
		return true
	} else if current+1 >= target {
		return true
	}
	return false
}

func (current mapNode) checkPath(heightMap [][]rune, visited map[string]int, step int) int {

	step++
	//right
	if inBounds(current.x+1, current.y, len(heightMap[current.y]), len(heightMap)) && visited[string(current.x+1)+" "+string(current.y)] > step { //in bounds and not same as parent
		if notTooHigh(current.value, heightMap[current.y][current.x+1]) { //TO DO they can move down inf number
			visited[string(current.x)+" "+string(current.y)] = step
			fmt.Println(string(current.value), current.x, current.y, "ok to move right")
			if heightMap[current.y][current.x+1] == 'E' {
				return step + 1
			}
			rightNode := mapNode{&current, current.x + 1, current.y, heightMap[current.y][current.x+1]}
			rightNode.checkPath(heightMap, visited, step)
			return -1

		}
	}
	//left
	if inBounds(current.x-1, current.y, len(heightMap[current.y]), len(heightMap)) && visited[string(current.x-1)+" "+string(current.y)] > step { //in bounds and not same as parent
		if notTooHigh(current.value, heightMap[current.y][current.x-1]) { //TO DO they can move down inf number
			visited[string(current.x)+" "+string(current.y)] = step
			fmt.Println(string(current.value), current.x, current.y, "ok to move right")
			if heightMap[current.y][current.x-1] == 'E' {
				return step + 1
			}
			rightNode := mapNode{&current, current.x - 1, current.y, heightMap[current.y][current.x-1]}
			rightNode.checkPath(heightMap, visited, step)
			return -1

		}
	}
	//down
	if inBounds(current.x, current.y+1, len(heightMap[current.y]), len(heightMap)) && visited[string(current.x)+" "+string(current.y+1)] > step { //in bounds and not same as parent
		if notTooHigh(current.value, heightMap[current.y+1][current.x]) { //TO DO they can move down inf number
			visited[string(current.x)+" "+string(current.y)] = step
			fmt.Println(string(current.value), current.x, current.y, "ok to move right")
			if heightMap[current.y+1][current.x] == 'E' {
				return step + 1
			}
			rightNode := mapNode{&current, current.x, current.y + 1, heightMap[current.y+1][current.x]}
			rightNode.checkPath(heightMap, visited, step)
			return -1

		}
	}
	//up
	if inBounds(current.x, current.y-1, len(heightMap[current.y]), len(heightMap)) && visited[string(current.x)+" "+string(current.y-1)] > step { //in bounds and not same as parent
		if notTooHigh(current.value, heightMap[current.y-1][current.x]) { //TO DO they can move down inf number
			visited[string(current.x)+" "+string(current.y)] = step
			fmt.Println(string(current.value), current.x, current.y, "ok to move right")
			if heightMap[current.y-1][current.x] == 'E' {
				return step + 1
			}
			rightNode := mapNode{&current, current.x, current.y - 1, heightMap[current.y-1][current.x]}
			rightNode.checkPath(heightMap, visited, step)
			return -1

		}
	}
	return -1
}

// if current.x-1 > -1 && (current.parent != nil || current.parent.x != current.x-1) && visited[string(current.x-1)+" "+string(current.y)] == 0 { //in bounds and not same as parent
// 	if Abs(int(current.value-heightMap[current.y][current.x-1])) < 2 { //
// 		fmt.Println(string(current.value), current.x, current.y, "ok to move left")
// 		leftNode := mapNode{&current, current.x - 1, current.y, heightMap[current.y][current.x-1]}
// 		leftNode.checkPath(heightMap, visited)
// 		//rightNode = mapNode{*current, coord{current.location.x+1}
// 	}
// }
// //down
// if current.y+1 < len(heightMap)-1 && (current.parent == nil || current.parent.y != current.y+1) && visited[string(current.x)+" "+string(current.y+1)] == 0 { //in bounds and not same as parent
// 	if Abs(int(current.value-heightMap[current.x][current.y+1])) < 2 { //
// 		fmt.Println(string(current.value), current.x, current.y, "ok to move down")
// 		downNode := mapNode{&current, current.x, current.y + 1, heightMap[current.y+1][current.x]}
// 		downNode.checkPath(heightMap, visited)
// 		//rightNode = mapNode{*current, coord{current.location.x+1}
// 	}
// }
// //up
// if current.y-1 > -1 && (current.parent != nil || current.parent.y != current.y-1) && visited[string(current.x)+" "+string(current.y-1)] == 0 { //in bounds and not same as parent
// 	if Abs(int(current.value-heightMap[current.x][current.y-1])) < 2 { //
// 		fmt.Println(string(current.value), current.x, current.y, "ok to move up")
// 		downNode := mapNode{&current, current.x, current.y + 1, heightMap[current.y+1][current.x]}
// 		downNode.checkPath(heightMap, visited)
// 		//rightNode = mapNode{*current, coord{current.location.x+1}
// 	}
// }

//determine shortest distance from items returned

// func checkPath(x int, y int, fromX int, fromY int) (length int) {
// 	var possiblePaths [4][4]int
// 	length := 1
// 	newY := y - 1
// 	//try up
// 	if newY < 0 || newY == fromY && x == fromX {
// 		//do nothing: new coord is out of bounds or parent direction
// 	} else if Abs(value-parent.Value) < 2 {
// 		//check path

// 	}
// 	else

// }
