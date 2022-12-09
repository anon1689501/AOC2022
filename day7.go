package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	inputText, err := os.Open("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	directoryValue := make(map[string]int)
	// directorySums := make(map[string]int)
	directoryContains := make(map[string][]string)

	currentDirectory := ""
	totalSum := 0

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		if input[0] == "$" && input[1] == "cd" { //save the full dir path in a string with split by ""
			if input[2] == ".." { //go up a directory
				currentDirectory = strings.Join(strings.Fields(currentDirectory)[:len(strings.Fields(currentDirectory))-1], " ")
			} else { //set up for the initial directory of /
				if currentDirectory == "" {
					currentDirectory = input[2]
				} else { //save folder name to currentDirectory
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
			for i := range strings.Fields(currentDirectory) {
				currentDirectorySlice := strings.Fields(currentDirectory)
				partialDirectory := strings.Join(strings.Fields(currentDirectory)[:len(currentDirectorySlice)-i], " ")
				// fmt.Println("full", currentDirectory, "partial", partialDirectory, "size", size)
				directoryValue[partialDirectory] += size
			}
		}
	}
	// fmt.Println("contains", directoryContains)
	// fmt.Println("values", directoryValue)

	// for mainDir, listOfSubDir := range directoryContains {
	// 	for _, fullDir := range listOfSubDir {
	// 		fmt.Println("main", mainDir)
	// 		for length, _ := range strings.Fields(fullDir) {
	// 			fmt.Println("sub", fullDir)
	// 			directorySums[mainDir] += directoryValue[strings.Join(strings.Fields(fullDir)[length:], " ")]
	// 			fmt.Println(mainDir, ":", fullDir, ",", strings.Join(strings.Fields(fullDir)[length:], "value"))
	// 		}
	// 		fmt.Println("")

	// 	}
	// }
	// fmt.Println(directorySums)

	for _, dirSum := range directoryValue {
		if dirSum <= 100000 {
			totalSum += dirSum
		}
	}

	fmt.Println("total sum", totalSum)

	//fmt.Println(directoryValue["/"])

	valueToDelete := 70000000 //also max size of disk
	spaceRequired := 30000000 - (valueToDelete - directoryValue["/"])

	for _, dirSum := range directoryValue {
		if dirSum < valueToDelete && dirSum > spaceRequired {
			valueToDelete = dirSum
		}
	}

	fmt.Println("folder value to delete:", valueToDelete)

}

// func iterateSubDir_copy(Sums map[string]int, Contains map[string][]string, directory string) (dirSum int) {
// 	dirSum = 0

// 	//fmt.Println("call with", directory)
// 	if _, subDirFound := Contains[directory]; subDirFound {
// 		for _, subDir := range Contains[directory] {
// 			dirSum += iterateSubDir_copy(Sums, Contains, subDir)
// 		}
// 	}
// 	return Sums[directory] + dirSum
// }

//1149424 too low
//1257239 too low
//1571315 too low

//47870454 possible total size

//8583399 too low
//8652144 too low
//8934862 too low
//9724718 uknown
