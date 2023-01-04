package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Day23() {

	inputText, err := os.Open("input/day23.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	y := 0

	elves := make(map[xy]bool)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		for x, v := range line {
			if v == '#' {
				elves[xy{x, y}] = true
			}
		}
		y++
	}
	fmt.Println(elves)
	printEvles(elves)
	//elfCount := len(elves)
	step := 0
	elvesNew := make(map[xy]bool)
	for {
		didNotMove := true
		elvesNew = make(map[xy]bool) //clear before every step
		for elf, _ := range elves {
			if orthoClear(elf, elves) {
				elvesNew[xy{elf.x, elf.y}] = true //does not move

			} else { //attempt to move
				didNotMove = false
				switch step % 4 {
				case 0:
					if topClear(elf, elves) {
						moveUp(elf, elvesNew)
					} else if bottomClear(elf, elves) {
						moveDown(elf, elvesNew)
					} else if leftClear(elf, elves) {
						moveLeft(elf, elvesNew)
					} else if rightClear(elf, elves) {
						moveRight(elf, elvesNew)
					} else { //cant move
						elvesNew[xy{elf.x, elf.y}] = true
					}
				case 1:
					if bottomClear(elf, elves) {
						moveDown(elf, elvesNew)
					} else if leftClear(elf, elves) {
						moveLeft(elf, elvesNew)
					} else if rightClear(elf, elves) {
						moveRight(elf, elvesNew)
					} else if topClear(elf, elves) {
						moveUp(elf, elvesNew)
					} else { //cant move
						elvesNew[xy{elf.x, elf.y}] = true
					}
				case 2:
					if leftClear(elf, elves) {
						moveLeft(elf, elvesNew)
					} else if rightClear(elf, elves) {
						moveRight(elf, elvesNew)
					} else if topClear(elf, elves) {
						moveUp(elf, elvesNew)
					} else if bottomClear(elf, elves) {
						moveDown(elf, elvesNew)
					} else { //cant move
						elvesNew[xy{elf.x, elf.y}] = true
					}

				case 3:
					if rightClear(elf, elves) {
						moveRight(elf, elvesNew)
					} else if topClear(elf, elves) {
						moveUp(elf, elvesNew)
					} else if bottomClear(elf, elves) {
						moveDown(elf, elvesNew)
					} else if leftClear(elf, elves) {
						moveLeft(elf, elvesNew)
					} else { //cant move
						elvesNew[xy{elf.x, elf.y}] = true
					}

				}
			}

		}
		elves = elvesNew
		step++
		if didNotMove { //step >= 10 {
			fmt.Println("did not move on step", step)
			break
		}

	}
	printEvles(elves)
}

func moveLeft(elf xy, elvesNew map[xy]bool) {
	if elvesNew[xy{elf.x - 1, elf.y}] { //someone already moved to that spot
		elvesNew[xy{elf.x, elf.y}] = true      //keep original spot
		elvesNew[xy{elf.x - 2, elf.y}] = true  //move the one that moved here back to its origial spot
		delete(elvesNew, xy{elf.x - 1, elf.y}) //delete that moved one
	} else {
		elvesNew[xy{elf.x - 1, elf.y}] = true
	}
}

func moveRight(elf xy, elvesNew map[xy]bool) {
	if elvesNew[xy{elf.x + 1, elf.y}] { //someone already moved to that spot
		elvesNew[xy{elf.x, elf.y}] = true      //keep original spot
		elvesNew[xy{elf.x + 2, elf.y}] = true  //move the one that moved here back to its origial spot
		delete(elvesNew, xy{elf.x + 1, elf.y}) //delete that moved one
	} else {
		elvesNew[xy{elf.x + 1, elf.y}] = true
	}
}

func moveUp(elf xy, elvesNew map[xy]bool) {
	if elvesNew[xy{elf.x, elf.y - 1}] { //someone already moved to that spot
		elvesNew[xy{elf.x, elf.y}] = true      //keep original spot
		elvesNew[xy{elf.x, elf.y - 2}] = true  //move the one that moved here back to its origial spot
		delete(elvesNew, xy{elf.x, elf.y - 1}) //delete that moved one
	} else {
		elvesNew[xy{elf.x, elf.y - 1}] = true
	}
}

func moveDown(elf xy, elvesNew map[xy]bool) {
	if elvesNew[xy{elf.x, elf.y + 1}] { //someone already moved to that spot
		elvesNew[xy{elf.x, elf.y}] = true      //keep original spot
		elvesNew[xy{elf.x, elf.y + 2}] = true  //move the one that moved here back to its origial spot
		delete(elvesNew, xy{elf.x, elf.y + 1}) //delete that moved one
	} else {
		elvesNew[xy{elf.x, elf.y + 1}] = true
	}
}

func leftClear(center xy, elves map[xy]bool) bool {
	if _, ok := elves[xy{center.x - 1, center.y + 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x - 1, center.y}]; ok {
		return false
	} else if _, ok := elves[xy{center.x - 1, center.y - 1}]; ok {
		return false
	}
	return true
}

func rightClear(center xy, elves map[xy]bool) bool {
	if _, ok := elves[xy{center.x + 1, center.y + 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x + 1, center.y}]; ok {
		return false
	} else if _, ok := elves[xy{center.x + 1, center.y - 1}]; ok {
		return false
	}
	return true
}

func topClear(center xy, elves map[xy]bool) bool {
	if _, ok := elves[xy{center.x - 1, center.y - 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x, center.y - 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x + 1, center.y - 1}]; ok {
		return false
	}
	return true
}

func bottomClear(center xy, elves map[xy]bool) bool {
	if _, ok := elves[xy{center.x - 1, center.y + 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x, center.y + 1}]; ok {
		return false
	} else if _, ok := elves[xy{center.x + 1, center.y + 1}]; ok {
		return false
	}
	return true
}

func orthoClear(center xy, elves map[xy]bool) bool {
	if leftClear(center, elves) && rightClear(center, elves) && topClear(center, elves) && bottomClear(center, elves) {
		return true
	}
	return false
}

func printEvles(elves map[xy]bool) {
	maxX := math.MinInt
	minX := math.MaxInt

	maxY := math.MinInt
	minY := math.MaxInt

	for location, _ := range elves {
		if location.x > maxX {
			maxX = location.x
		}
		if location.x < minX {
			minX = location.x
		}
		if location.y > maxY {
			maxY = location.y
		}
		if location.y < minY {
			minY = location.y
		}
	}
	fmt.Println("min x", minX, "max x", maxX, "min y", minY, "max y", maxY)
	area := (maxX-minX+1)*(maxY-minY+1) - len(elves)

	fmt.Println(area)

}
