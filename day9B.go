package main

import (
	"bufio"
	"fmt"
	"log"

	"os"
	"strconv"
	"strings"
)

func Day9B() {
	inputText, err := os.Open("input/day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	visited := make(map[string]int)

	var knots []Position

	for i := 0; i < len(knots); i++ {
		knots[i].xCoord = 0
		knots[i].yCoord = 0
	}

	visited["0 0"] = 1

	//writing file area

	// f, err := os.OpenFile("outputB.txt", os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	//end of write file area
	inputLine := 0
	scanner := bufio.NewScanner(inputText)
	for scanner.Scan() {

		input := strings.Fields(scanner.Text())
		direction := input[0]
		//fmt.Println(direction)
		distance, _ := strconv.Atoi(input[1])
		//fmt.Println(distance)
		//fmt.Println(knots[len(knots)-1])

		switch direction {

		case "L":
			//fmt.Print(direction, distance)
			for i := 0; i < distance; i++ {
				//fmt.Print(knots[0])
				knots[0].xCoord--
				// fmt.Println("head", inputLine, knots[0], direction, distance)
				for j := 0; j < len(knots)-1; j++ {
					xMove, yMove := move(knots[j], knots[j+1])
					knots[j+1].xCoord += xMove
					knots[j+1].yCoord += yMove

					if j+1 == len(knots)-1 {
						visited[strconv.Itoa(knots[len(knots)-1].xCoord)+" "+strconv.Itoa(knots[len(knots)-1].yCoord)]++
					}

				}
				//fmt.Println(knots)

			}

		case "R":
			//fmt.Print(direction, distance)
			for i := 0; i < distance; i++ {
				//fmt.Print(knots[0])
				knots[0].xCoord++
				// fmt.Println("head", inputLine, knots[0], direction, distance)

				for j := 0; j < len(knots)-1; j++ {
					xMove, yMove := move(knots[j], knots[j+1])
					knots[j+1].xCoord += xMove
					knots[j+1].yCoord += yMove
					if j+1 == len(knots)-1 {
						visited[strconv.Itoa(knots[len(knots)-1].xCoord)+" "+strconv.Itoa(knots[len(knots)-1].yCoord)]++
					}
				}
				//fmt.Println(knots)

			}

		case "D":
			//fmt.Print(direction, distance)
			for i := 0; i < distance; i++ {
				//fmt.Print(knots[0])
				knots[0].yCoord--
				// fmt.Println("head", inputLine, knots[0], direction, distance)

				for j := 0; j < len(knots)-1; j++ {
					xMove, yMove := move(knots[j], knots[j+1])
					knots[j+1].xCoord += xMove
					knots[j+1].yCoord += yMove

					if j+1 == len(knots)-1 {
						visited[strconv.Itoa(knots[len(knots)-1].xCoord)+" "+strconv.Itoa(knots[len(knots)-1].yCoord)]++
						//fmt.Println(inputLine, knots[8], direction, distance, knots[len(knots)-1])
					}

				}
				//fmt.Println(knots)

			}

		case "U":
			//fmt.Print(direction, distance)
			for i := 0; i < distance; i++ {
				//fmt.Print(knots[0])
				knots[0].yCoord++
				// fmt.Println("head", inputLine, knots[0], direction, distance)

				for j := 0; j < len(knots)-1; j++ {
					xMove, yMove := move(knots[j], knots[j+1])
					knots[j+1].xCoord += xMove
					knots[j+1].yCoord += yMove

					if j+1 == len(knots)-1 {
						visited[strconv.Itoa(knots[len(knots)-1].xCoord)+" "+strconv.Itoa(knots[len(knots)-1].yCoord)]++
						//fmt.Println(inputLine, knots[8], direction, distance, knots[len(knots)-1])
					}

				}
				//fmt.Println(knots)

			}

		default:
			fmt.Println("unknown direction")
		}
		inputLine++
	}
	fmt.Println(len(visited))
	//1033 too low
	//1034 too low
	//1224 too low
}
