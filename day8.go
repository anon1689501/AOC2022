package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day8() {
	inputText, err := os.Open("input/day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	// trees := make([][]int, 0)
	var trees [99][99]int

	treeRowCount := 0
	for scanner.Scan() {
		for i, treeHeight := range scanner.Text() {
			//fmt.Println(i)
			trees[treeRowCount][i], _ = strconv.Atoi(string(treeHeight))
		}
		treeRowCount++
	}
	//fmt.Println(trees)
	treeVisible := make(map[string]int)
	for row := 1; row < 98; row++ { //dont need do row 0 or do 98 because its an edge
		tallestLeft := trees[row][0]
		//fmt.Println(tallestLeft)
		for column := 1; column < 99; column++ {
			if trees[row][column] > tallestLeft {

				treeVisible[strconv.Itoa(row)+" "+strconv.Itoa(column)] += trees[row][column]
				tallestLeft = trees[row][column]
			}
		}
	}
	for row := 1; row < 98; row++ {
		tallestRight := trees[row][98]
		//fmt.Println(tallestRight)
		for column := 97; column > 0; column-- {
			if trees[row][column] > tallestRight {

				treeVisible[strconv.Itoa(row)+" "+strconv.Itoa(column)] += trees[row][column]
				tallestRight = trees[row][column]
			}
		}
	}
	for column := 1; column < 98; column++ {
		tallestTop := trees[0][column]
		//fmt.Println(tallestTop)
		//fmt.Println(tallestLeft)
		for row := 1; row < 99; row++ {
			if trees[row][column] > tallestTop {
				if column == 0 {
					fmt.Println(row, trees[row][column])
				}
				treeVisible[strconv.Itoa(row)+" "+strconv.Itoa(column)] += trees[row][column]
				tallestTop = trees[row][column]
			}
		}
	}
	for column := 1; column < 98; column++ {
		tallestBottom := trees[98][column]
		//fmt.Println(tallestRight)
		for row := 97; row > 0; row-- {
			if trees[row][column] > tallestBottom {
				treeVisible[strconv.Itoa(row)+" "+strconv.Itoa(column)] += trees[row][column]
				tallestBottom = trees[row][column]
			}
		}
	}

	fmt.Println(len(treeVisible) + 392)
	// x := 2
	// y := 6
	// fmt.Println(x, ",", y, ":", trees[x][y], "right:", rightView(x, y, trees), "bottom:", bottomView(x, y, trees), "top:", topView(x, y, trees), "left:", leftView(x, y, trees))
	highestScenic := 0
	for row := 1; row < 98; row++ {
		for column := 1; column < 98; column++ {
			scenic := bottomView(row, column, trees) * topView(row, column, trees) * rightView(row, column, trees) * leftView(row, column, trees)
			if scenic > highestScenic {
				highestScenic = scenic
			}

		}
	}
	fmt.Println(highestScenic)
}

// func leftView(row int, column int, trees [][]int) (distance int){
// 	for
// }

func bottomView(r int, c int, trees [99][99]int) (distance int) {
	distance = 0
	for row := r + 1; row < 99; row++ {
		if trees[r][c] > trees[row][c] {
			distance++
		} else {
			distance++
			return distance
		}
	}
	return distance
}

func topView(r int, c int, trees [99][99]int) (distance int) {
	distance = 0
	for row := r - 1; row >= 0; row-- {
		if trees[r][c] > trees[row][c] {
			distance++
		} else {
			distance++
			return distance
		}
	}
	return distance
}

func rightView(r int, c int, trees [99][99]int) (distance int) {
	distance = 0
	for column := c + 1; column < 99; column++ {
		if trees[r][c] > trees[r][column] {
			distance++
		} else {
			distance++
			return distance
		}
	}

	return distance
}

func leftView(r int, c int, trees [99][99]int) (distance int) {
	distance = 0
	for column := c - 1; column >= 0; column-- {
		if trees[r][c] > trees[r][column] {
			distance++
		} else {
			distance++
			return distance
		}
	}

	return distance
}

//1541 too low
//1815 too high
