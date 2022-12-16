package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day15() {

	inputText, err := os.Open("input/day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	//cave := make(map[xy]bool)
	//sensorDistance := make(map[xy]int)

	scanner := bufio.NewScanner(inputText)

	// minX := 10000000
	// maxX := 0

	startLine := 0
	endLine := 4000000

	//targetRow := 2000000
	//rowDetails := make(map[int]rune)
	pairs := make([]pair, 0)
	sensors := make([]sensor, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "x=", "")
		line = strings.ReplaceAll(line, "y=", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ":", "")
		//fmt.Println(line)
		fields := strings.Fields(line)
		sensorX, _ := strconv.Atoi(fields[2])
		sensorY, _ := strconv.Atoi(fields[3])
		beaconX, _ := strconv.Atoi(fields[8])
		beaconY, _ := strconv.Atoi(fields[9])
		sensors = append(sensors, sensor{sensorX, sensorY, findDistance(sensorX, sensorY, beaconX, beaconY)})

	}

	for line := startLine; line <= endLine; line++ {
		pairs = nil
		for _, pairVal := range sensors {
			tempPair := singleLine(pairVal, line)
			if tempPair.max == 0 && tempPair.min == 0 {
				continue
			}
			pairs = append(pairs, tempPair)
			//fmt.Println(pairs)

		}
		//fmt.Println(pairs)
		pairs = sortPairs(pairs)
		//fmt.Println(pairs)
		//front := pairs[0].min
		end := pairs[0].max
		for i := 1; i < len(pairs); i++ {
			if pairs[i].min <= end { //ok
				if pairs[i].max > end {
					end = pairs[i].max
				}

			} else if end+2 == pairs[i].min {
				fmt.Println("skipped one at row:", line, "column:", end+1, (end+1)*4000000+line)

			} else {
				//fmt.Println("row:", line, "column:", pairs[i].min)
			}
		}
		//fmt.Println(front, end)

	}

	// fmt.Println(pairs)
	//start := pairs[0].min
	// end := pairs[0].max
	// for i := 1; i < len(pairs)-1; i++ {
	// 	if pairs[i].min <= end { //ok
	// 		if pairs[i].max > end {
	// 			end = pairs[i].max
	// 		}

	// 	} else {
	// 		fmt.Println("row:", l, "column:", pairs[i].min)
	// 		break
	// 	}
	// }

	//d15Print(cave)
	// nonBeaconCount := 0

	// for i, v := range rowDetails {
	// 	if i == row && v == '#' {
	// 		nonBeaconCount++
	// 	}

	// }
	//fmt.Println(len(rowDetails))
	// fmt.Println(nonBeaconCount)
	//fmt.Println(maxX - minX)
}

func sortPairs(pairs []pair) []pair {
	for i := 0; i < len(pairs)-1; i++ {
		for j := 0; j < len(pairs)-i-1; j++ {
			//fmt.Println(linesB[j], linesB[j+1])
			if !pairOrdered(pairs[j].min, pairs[j+1].min) {
				pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			}
		}
	}
	return pairs

}

func pairOrdered(left int, right int) bool {
	return left <= right
}

type pair struct {
	min int
	max int
}

type sensor struct {
	x        int
	y        int
	distance int
}

func findDistance(sensorX int, sensorY int, beaconX int, beaconY int) int {
	return Abs(sensorX-beaconX) + Abs(sensorY-beaconY)
}

func d15Print(cave map[xy]rune) {
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

func mapSensor(sensorX int, sensorY int, distance int, cave map[xy]bool) { //cave map[xy]rune,
	for i := 0; i <= distance; i++ {
		for j := 0; j <= distance-i; j++ {
			if sensorX+j < 0 || sensorX+j > 4000000 {
				continue
			}
			if sensorX-j < 0 || sensorX-j > 4000000 {
				continue
			}
			if sensorY+j < 0 || sensorY+j > 4000000 {
				continue
			}
			if sensorY-j < 0 || sensorY-j > 4000000 {
				continue
			}
			if !cave[xy{sensorX + j, sensorY + i}] { //bottom right
				cave[xy{sensorX + j, sensorY + i}] = true
			}
			if !cave[xy{sensorX - j, sensorY - i}] { //top left
				cave[xy{sensorX - j, sensorY - i}] = true
			}
			if !cave[xy{sensorX + j, sensorY - i}] { //top right
				cave[xy{sensorX + j, sensorY - i}] = true
			}
			if !cave[xy{sensorX - j, sensorY + i}] { //top right
				cave[xy{sensorX - j, sensorY + i}] = true
			}
		}
	}

}

// func mapSensor(sensorX int, sensorY int, distance int, row int, rowDetails map[int]rune) { //cave map[xy]rune,
// 	for i := 0; i <= distance; i++ {
// 		for j := 0; j <= distance-i; j++ {
// 			if sensorY+i == row && rowDetails[sensorX+j] == 0 { //bottom right
// 				rowDetails[sensorX+j] = '#'
// 			}
// 			if sensorY-i == row && rowDetails[sensorX-j] == 0 { //top left
// 				rowDetails[sensorX-j] = '#'
// 			}
// 			if sensorY-i == row && rowDetails[sensorX+j] == 0 { //top right
// 				rowDetails[sensorX+j] = '#'
// 			}
// 			if sensorY+i == row && rowDetails[sensorX-j] == 0 { //top right
// 				rowDetails[sensorX-j] = '#'
// 			}
// 		}
// 	}

// }

func singleLine(sen sensor, row int) (myPair pair) { //cave map[xy]rune,

	myPair.min = 0
	myPair.max = 0

	if sen.y < row && sen.y+sen.distance > row { //sensor above (lower y value ) than target row
		remainingDistance := sen.distance - Abs(row-sen.y) //row 10 distance 20 target 20
		myPair.min = sen.x - remainingDistance
		myPair.max = sen.x + remainingDistance

	} else if sen.y > row && sen.y-sen.distance < row {
		remainingDistance := sen.distance - Abs(row-sen.y) //row 10 distance 20 target 20
		myPair.min = sen.x - remainingDistance
		myPair.max = sen.x + remainingDistance

	}
	return myPair

}
