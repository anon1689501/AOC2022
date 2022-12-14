package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day14() {

	inputText, err := os.Open("input/day14.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	cave := make(map[xy]rune)
	//cave[xy{500, 0}] = '+'

	deepest := 0

	for scanner.Scan() {
		lines := make([]string, 0)
		line := scanner.Text()
		//fmt.Println(line)
		line = strings.ReplaceAll(line, "->", ",")
		//fmt.Println(line)
		line = strings.ReplaceAll(line, " ", "")
		//fmt.Println("no whitespace", line)
		lines = strings.Split(line, ",")
		//fmt.Println(lines)
		for i := 0; i < len(lines)-2; i += 2 {
			x1, _ := strconv.Atoi(lines[i])
			y1, _ := strconv.Atoi(lines[i+1])
			x2, _ := strconv.Atoi(lines[i+2])
			y2, _ := strconv.Atoi(lines[i+3])
			//fmt.Println(x1, y1, x2, y2)
			if x1 == x2 {
				buildVert(x1, y1, y2, cave)
			} else {
				buildHori(y1, x1, x2, cave)
			}

			if y1 > deepest {
				deepest = y1
			}
			if y1 > deepest {
				deepest = y2
			}

		}

	}
	//fmt.Println(cave)
	//do sand
	sandLanded := true
	//fmt.Println(deepest)

	for i := 0; ; i++ {
		sandLanded = doSand(500, 0, cave, deepest+2)

		if !sandLanded || cave[xy{500, 0}] != 0 {
			cave[xy{500, 0}] = '+'
			fmt.Println(i + 1)
			break
		}
	}

	cavePrint(cave)

}
func doSand(x int, y int, cave map[xy]rune, limit int) bool {
	for i := y; i <= limit; i++ {
		//fmt.Println("trying", x, i, cave[xy{x, i}])
		if i == limit {
			//fmt.Println("limit reached", x, i)
			cave[xy{x, i}] = '#'
			cave[xy{x + 1, i}] = '#'
			cave[xy{x - 1, i}] = '#'
		}
		if cave[xy{x, i}] != 0 { //find vertial clear point
			if cave[xy{x - 1, i}] != 0 && cave[xy{x + 1, i}] != 0 { //if bot left and bot right are blocked
				cave[xy{x, i - 1}] = 'o'
				return true
			} else if cave[xy{x - 1, i}] == 0 { //&& cave[xy{x - 1, i + 1}] != 0 { //if bottom left is open and the one right under that is blocked
				return doSand(x-1, i, cave, limit)
			} else if cave[xy{x + 1, i}] == 0 { //&& cave[xy{x + 1, i + 1}] != 0 { //if bottom right is open and the one right under that is blocked
				return doSand(x+1, i, cave, limit)
			}
			// if cave[xy{x - 1, i + 1}] == 0 { //test bottom left
			// 	if doSand(x-1, i+1, cave) { //try to drop vertically from bottom left
			// 		return true //if placed in the recursive call
			// 	} else {
			// 		cave[xy{x - 1, i + 1}] == 'o'
			// 		return true
			// 	}
			// }

			//check bot right
		}
	}
	return false

}
func cavePrint(cave map[xy]rune) {
	minX := 1000
	maxX := 0
	minY := 1000
	maxY := 0
	for i, _ := range cave {
		if i.x > maxX {
			maxX = i.x
		}
		if i.x < minX {
			minX = i.x
		}
		if i.y > maxY {
			maxY = i.y
		}
		if i.y < minY {
			minY = i.y
		}
	}
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if cave[xy{j, i}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(cave[xy{j, i}]))
			}
		}
		fmt.Println(" ")
	}

}
func buildHori(y int, x1 int, x2 int, cave map[xy]rune) {

	if x1 > x2 { //up and -1
		for i := x1; i >= x2; i-- {
			cave[xy{i, y}] = '#'
		}
	} else {
		for i := x1; i <= x2; i++ {
			cave[xy{i, y}] = '#'
		}
	}

}
func buildVert(x int, y1 int, y2 int, cave map[xy]rune) {

	if y1 > y2 { //up and -1
		for i := y1; i >= y2; i-- {
			cave[xy{x, i}] = '#'
		}
	} else {
		for i := y1; i <= y2; i++ {
			cave[xy{x, i}] = '#'
		}
	}

}

type xy struct {
	x int
	y int
}
