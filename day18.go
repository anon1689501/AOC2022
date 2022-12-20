package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day18() {

	inputText, err := os.Open("input/day18.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()
	cubes := make([]cube, 0)

	scanner := bufio.NewScanner(inputText)

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		localX, _ := strconv.Atoi(coords[0])
		localY, _ := strconv.Atoi(coords[1])
		localZ, _ := strconv.Atoi(coords[2])
		cubes = append(cubes, cube{localX, localY, localZ})
	}

	matchingFaces := 0
	for i := 0; i < len(cubes)-1; i++ {
		for j := i + 1; j < len(cubes); j++ {
			//fmt.Println(linesB[j], linesB[j+1])
			if cubes[i].x == cubes[j].x && cubes[i].y == cubes[j].y && Abs(cubes[i].z-cubes[j].z) < 2 {
				matchingFaces++
			} else if cubes[i].y == cubes[j].y && cubes[i].z == cubes[j].z && Abs(cubes[i].x-cubes[j].x) < 2 {
				matchingFaces++
			} else if cubes[i].z == cubes[j].z && cubes[i].x == cubes[j].x && Abs(cubes[i].y-cubes[j].y) < 2 {
				matchingFaces++
			}
		}
	}
	fmt.Println("Part 1 matching:", matchingFaces, "total faces:", len(cubes)*6, "surface area:", len(cubes)*6-matchingFaces*2)

	missingCubes := make(map[cube]int)
	cubes = cubeSort("x", cubeSort("y", cubeSort("z", cubes)))
	//fmt.Println(cubes)

	compareX := -1
	compareY := -1
	compareZ := -1
	for _, c := range cubes {
		if compareZ == -1 { // new x,y to look at
			compareX = c.x
			compareY = c.y
			compareZ = c.z
		} else if c.x == compareX && c.y == compareY { //x,y match previous iteration
			if !(compareZ+1 == c.z) {
				for i := compareZ + 1; i < c.z; i++ {
					missingCubes[cube{c.x, c.y, i}]++
					//fmt.Println("added missing: ", compareX, compareY, i, "val:", missingCubes[cube{c.x, c.y, i}])
				}
			}
			//} else { //if x y do not match reset everything for next set of x,y

		}
		compareX = c.x
		compareY = c.y
		compareZ = c.z
	}

	cubes = cubeSort("z", cubeSort("x", cubeSort("y", cubes)))
	//fmt.Println(cubes)
	compareX = -1
	compareY = -1
	compareZ = -1
	for _, c := range cubes {
		if compareZ == -1 { // new x,y to look at
			compareX = c.x
			compareY = c.y
			compareZ = c.z
		} else if c.z == compareZ && c.x == compareX { //x,y match previous iteration
			if !(compareY+1 == c.y) {
				for i := compareY + 1; i < c.y; i++ {
					missingCubes[cube{c.x, i, c.z}]++
					//fmt.Println("added missing: ", compareX, i, compareZ, "val:", missingCubes[cube{c.x, i, c.z}])
				}
			}
			//} else { //if x y do not match reset everything for next set of x,y

		}
		compareX = c.x
		compareY = c.y
		compareZ = c.z
	}

	cubes = cubeSort("y", cubeSort("z", cubeSort("x", cubes)))
	//fmt.Println(cubes)
	compareX = -1
	compareY = -1
	compareZ = -1
	for _, c := range cubes {
		if compareZ == -1 { // new x,y to look at
			compareX = c.x
			compareY = c.y
			compareZ = c.z
		} else if c.y == compareY && c.z == compareZ { //x,y match previous iteration
			if !(compareX+1 == c.x) {
				for i := compareX + 1; i < c.x; i++ {
					missingCubes[cube{i, c.y, c.z}]++
					fmt.Println("added missing: ", i, compareY, compareZ, "val:", missingCubes[cube{i, c.y, c.z}])
				}
			}
			//} else { //if x y do not match reset everything for next set of x,y

		}
		compareX = c.x
		compareY = c.y
		compareZ = c.z
	}

	notMissingCount := countIfLessThan(missingCubes, 3)
	fmt.Println(notMissingCount)
	for {
		//fmt.Println(missingCubes)
		for cubeOuter, valOuter := range missingCubes {
			if valOuter < 3 {
				for cubeInner, valInner := range missingCubes {
					if valInner == 3 {
						if cubeOuter.x == cubeInner.x && cubeOuter.y == cubeInner.y && Abs(cubeOuter.z-cubeInner.z) < 2 {
							missingCubes[cubeInner] = 2
							fmt.Println("inner changed", cubeInner)
						} else if cubeOuter.y == cubeInner.y && cubeOuter.z == cubeInner.z && Abs(cubeOuter.x-cubeInner.x) < 2 {
							missingCubes[cubeInner] = 2
							fmt.Println("inner changed", cubeInner)
						} else if cubeOuter.z == cubeInner.z && cubeOuter.x == cubeInner.x && Abs(cubeOuter.y-cubeInner.y) < 2 {
							missingCubes[cubeInner] = 2
							fmt.Println("inner changed", cubeInner)
						}
					}
				}
			}
		}
		updatedNotMissingCount := countIfLessThan(missingCubes, 3)
		fmt.Println(updatedNotMissingCount)
		if notMissingCount < updatedNotMissingCount {
			fmt.Println("restart")
			notMissingCount = updatedNotMissingCount
		} else {
			fmt.Println("break")
			break
		}
	}

	confirmedMissingCube := make([]cube, 0)
	for i, v := range missingCubes {
		if v == 3 {
			confirmedMissingCube = append(confirmedMissingCube, i)
			//fmt.Println(i)

		}
	}

	matchingFacesMissing := 0
	for i := 0; i < len(confirmedMissingCube)-1; i++ {
		for j := i + 1; j < len(confirmedMissingCube); j++ {
			//fmt.Println(linesB[j], linesB[j+1])
			if confirmedMissingCube[i].x == confirmedMissingCube[j].x && confirmedMissingCube[i].y == confirmedMissingCube[j].y && Abs(confirmedMissingCube[i].z-confirmedMissingCube[j].z) < 2 {
				matchingFacesMissing++
			} else if confirmedMissingCube[i].y == confirmedMissingCube[j].y && confirmedMissingCube[i].z == confirmedMissingCube[j].z && Abs(confirmedMissingCube[i].x-confirmedMissingCube[j].x) < 2 {
				matchingFacesMissing++
			} else if confirmedMissingCube[i].z == confirmedMissingCube[j].z && confirmedMissingCube[i].x == confirmedMissingCube[j].x && Abs(confirmedMissingCube[i].y-confirmedMissingCube[j].y) < 2 {
				matchingFacesMissing++
			}
		}
	}
	insideFaces := len(confirmedMissingCube)*6 - matchingFacesMissing
	fmt.Println(insideFaces)

	fmt.Println("Part 2:", len(cubes)*6-matchingFaces*2-insideFaces)

}

func countIfLessThan(missingCubes map[cube]int, lessThanVal int) (count int) {
	count = 0
	for _, v := range missingCubes {
		if v < lessThanVal {
			count++
		}
	}
	return count
}

type cube struct {
	x int
	y int
	z int
}

func cubeSort(dim string, cubes []cube) []cube {
	if dim == "x" {
		for i := 0; i < len(cubes)-1; i++ {
			for j := 0; j < len(cubes)-i-1; j++ {
				//fmt.Println(linesB[j], linesB[j+1])
				if !pairOrdered(cubes[j].x, cubes[j+1].x) {
					cubes[j], cubes[j+1] = cubes[j+1], cubes[j]
				}
			}
		}

	} else if dim == "y" {
		for i := 0; i < len(cubes)-1; i++ {
			for j := 0; j < len(cubes)-i-1; j++ {
				//fmt.Println(linesB[j], linesB[j+1])
				if !pairOrdered(cubes[j].y, cubes[j+1].y) {
					cubes[j], cubes[j+1] = cubes[j+1], cubes[j]
				}
			}
		}

	} else if dim == "z" {
		for i := 0; i < len(cubes)-1; i++ {
			for j := 0; j < len(cubes)-i-1; j++ {
				//fmt.Println(linesB[j], linesB[j+1])
				if !pairOrdered(cubes[j].z, cubes[j+1].z) {
					cubes[j], cubes[j+1] = cubes[j+1], cubes[j]
				}
			}
		}

	}
	return cubes
} //part 2 1038 too low
