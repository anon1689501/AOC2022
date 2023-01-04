package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day22() {

	inputText, err := os.Open("input/day22.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	jungleMap := make([][]rune, 0)

	scanner := bufio.NewScanner(inputText)

	//column := 0
	blankFound := false
	moves := ""
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			blankFound = true
		}
		if !blankFound {
			runeRow := []rune(scanner.Text())
			jungleMap = append(jungleMap, runeRow)
		}
		if blankFound && len(scanner.Text()) > 0 {
			moves = scanner.Text()

		}

	}
	printMap(jungleMap)

	moveList := processMoves(moves)

	facing := 0 //0 is right, 1 is down, 2 is left, 3 is up
	x := 0
	y := 0

	for i := 0; i < len(moveList); i++ {
		currentMove := moveList[i]
		if currentMove == "R" {
			facing++
			facing %= 4
		} else if currentMove == "L" {
			facing += 3
			facing %= 4
		} else { //its a number
			distance, _ := strconv.Atoi(currentMove)
			x, y = move22(jungleMap, distance, facing, x, y)

		}
	}
}

func printMap(jungleMap [][]rune) {
	for _, row := range jungleMap {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println("")
	}

}

func move22(jungleMap [][]rune, distance int, facing int, x int, y int) (newX, newY int) {
	switch facing {
	case 0: //right
	complete:
		for i := 0; i < distance; {
			if jungleMap[y][x+i] == '#' { //stop at #
				jungleMap[y][x+i-1] = '>'
				x = x + i - 1
				break
			} else if jungleMap[y][x+i] != ' ' {
				jungleMap[y][x+i-1] = '>'
				i++
				continue
			} else if x+1 > len(jungleMap[y]) { //out of bounds at the end of the slice
				for j := 0; j < len(jungleMap[y]); j++ { //iterate through the start of the slice to find where the map starts
					if jungleMap[y][j] == '#' {
						x = len(jungleMap[y]) - 1
						break complete
					} else if jungleMap[y][j] == '.' {
						jungleMap[y][x+i-1] = '>' //this is a copy from the line above and will not work correctly when it wraps around TO DO, consider tracking the distance traveled separately from x
						x = j
						i++
						break
					}
				}
			}
			if i == distance {
				x += i
				break
			}
		}
	case 1: //down

	case 2: //left
		for i := 0; i < distance; i++ {
			if jungleMap[y][x-i] == '#' {
				x = x - i + 1
				break
			} else if x-1 < 0 {
				//for
			}
		}
	case 3: //up

	}
	newX = x
	newY = y
	return

}

func processMoves(input string) (output []string) {
	numberHolder := ""
	for _, v := range input {
		if v >= '0' && v <= '9' {
			numberHolder += string(v)
		} else {
			output = append(output, numberHolder)
			output = append(output, string(v))
			numberHolder = ""
		}
	}
	output = append(output, numberHolder) //assuming that the last command is a number

	return output
}
