package main

import (
	"bufio"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"
)

type Position struct {
	xCoord int
	yCoord int
}

func Day9() {
	inputText, err := os.Open("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	visited := make(map[string]int)
	head := Position{0, 0}
	tail := Position{0, 0}
	inputLine := 0

	visited["0 0"] = 1

	//file writing area

	// f, err := os.OpenFile("outputA.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	//end file writing area

	scanner := bufio.NewScanner(inputText)
	for scanner.Scan() {

		input := strings.Fields(scanner.Text())
		direction := input[0]
		//fmt.Println(direction)
		distance, _ := strconv.Atoi(input[1])
		//fmt.Println(distance)

		switch direction {

		case "L":
			for i := 0; i < distance; i++ {
				//fmt.Print(head)
				head.xCoord--
				if touching(head, tail) {
					//do nothing
				} else if head.yCoord == tail.yCoord {
					tail.xCoord--
				} else {
					tail.xCoord--
					tail.yCoord = head.yCoord
				}
				visited[strconv.Itoa(tail.xCoord)+" "+strconv.Itoa(tail.yCoord)]++
				//fmt.Fprintln(f, inputLine, head, direction, distance, tail)

			}

		case "R":
			for i := 0; i < distance; i++ {
				//fmt.Print(head)
				head.xCoord++
				if touching(head, tail) {
					//do nothing
				} else if head.yCoord == tail.yCoord {
					tail.xCoord++
				} else {
					tail.xCoord++
					tail.yCoord = head.yCoord
				}
				visited[strconv.Itoa(tail.xCoord)+" "+strconv.Itoa(tail.yCoord)]++
				//fmt.Fprintln(f, inputLine, head, direction, distance, tail)

			}

		case "D":
			for i := 0; i < distance; i++ {
				//fmt.Print(head)
				head.yCoord--
				if touching(head, tail) {
					//do nothing
				} else if head.xCoord == tail.xCoord {
					tail.yCoord--
				} else {
					tail.yCoord--
					tail.xCoord = head.xCoord
				}
				visited[strconv.Itoa(tail.xCoord)+" "+strconv.Itoa(tail.yCoord)]++
				//fmt.Fprintln(f, inputLine, head, direction, distance, tail)

			}

		case "U":
			for i := 0; i < distance; i++ {
				//fmt.Print(head)
				head.yCoord++
				if touching(head, tail) {
					//do nothing
				} else if head.xCoord == tail.xCoord {
					tail.yCoord++
				} else {
					tail.yCoord++
					tail.xCoord = head.xCoord
				}
				visited[strconv.Itoa(tail.xCoord)+" "+strconv.Itoa(tail.yCoord)]++
				//fmt.Fprintln(f, inputLine, head, direction, distance, tail)
			}

		default:
			fmt.Println("unknown direction")
		}
		inputLine++
	}
	fmt.Println(len(visited))
}

func touching(head Position, tail Position) bool { //logic probably could be simplified
	if Abs(head.xCoord-tail.xCoord) <= 1 && head.yCoord == tail.yCoord {
		return true //if col off by 1 and row is the same
	}
	if Abs(head.yCoord-tail.yCoord) <= 1 && head.xCoord == tail.xCoord {
		return true //if row off by 1 and col is the same
	}
	if Abs(head.xCoord-tail.xCoord) <= 1 && Abs(head.yCoord-tail.yCoord) <= 1 {
		return true //if diagonal
	}
	return false

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(head Position, tail Position) (xMove int, yMove int) { //logic probably could be simplified
	if touching(head, tail) {
		return 0, 0
	} else {
		if Abs(head.xCoord-tail.xCoord) > 1 && head.yCoord == tail.yCoord {
			if head.xCoord > tail.xCoord {
				return 1, 0
			} else {
				return -1, 0
			}
		}
		if Abs(head.yCoord-tail.yCoord) > 1 && head.xCoord == tail.xCoord {
			if head.yCoord > tail.yCoord {
				return 0, 1
			} else {
				return 0, -1
			}
		}
		if Abs(head.xCoord-tail.xCoord) >= 1 && Abs(head.yCoord-tail.yCoord) >= 1 {
			localX := 0
			localY := 0
			if head.xCoord > tail.xCoord {
				localX = 1
			} else {
				localX = -1
			}
			if head.yCoord > tail.yCoord {
				localY = 1
			} else {
				localY = -1
			}
			return localX, localY
		}
	}
	return 0, 0

}
