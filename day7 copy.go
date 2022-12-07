package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day7_copy() {
	inputText, err := os.Open("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	directorySums := make(map[string]int)
	directoryContains := make(map[string][]string)

	currentDirectory := ""
	totalSum := 0

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		if input[0] == "$" && input[1] == "cd" {
			if input[2] == ".." {
				currentDirectory = strings.Join(strings.Fields(currentDirectory)[:len(strings.Fields(currentDirectory))-1], " ")
			} else {
				if currentDirectory == "" {
					currentDirectory = input[2]
				} else {
					currentDirectory += " " + input[2]
				}
			}
		} else if input[0] == "$" {
			//probably ignore because its a ls
		} else if input[0] == "dir" {
			//do mapping
			directoryContains[currentDirectory] = append(directoryContains[currentDirectory], currentDirectory+" "+input[1])
		} else {
			//its a number
			size, _ := strconv.Atoi(input[0])
			directorySums[currentDirectory] += size
		}
	}
	// fmt.Println(directoryContains)
	// fmt.Println(directorySums)

	for mainDir, listOfSubDir := range directoryContains {
		for _, dir := range listOfSubDir {
			directorySums[mainDir] += iterateSubDir_copy(directorySums, directoryContains, dir)
		}
	}

	for _, dirSum := range directorySums {
		if dirSum <= 100000 {
			totalSum += dirSum
		}
	}
	fmt.Println("total sum", totalSum)

	fmt.Println(directorySums["/"])

}

func iterateSubDir_copy(Sums map[string]int, Contains map[string][]string, directory string) (dirSum int) {
	dirSum = 0

	//fmt.Println("call with", directory)
	if _, subDirFound := Contains[directory]; subDirFound {
		for _, subDir := range Contains[directory] {
			dirSum += iterateSubDir_copy(Sums, Contains, subDir)
		}
	}
	return Sums[directory] + dirSum
}

//1149424 too low
//1257239 too low
//1571315 too low

//47870454 possible total size
