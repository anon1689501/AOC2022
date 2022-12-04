package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day3() {
	//read input

	inputText, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)
	//runeArray := make([]rune, 1)
	sum := 0

	stringArray := make([]string, 0)

	for scanner.Scan() {
		mylen := len(scanner.Text()) / 2
		left := scanner.Text()[:mylen]
		right := scanner.Text()[mylen:]
		//fmt.Println(scanner.Text(), left, right)
		stringArray = append(stringArray, scanner.Text())

	out:
		for _, v := range left {
			for _, c := range right {
				if v == c {
					//runeArray = append(runeArray, v)
					if v < 96 {
						sum += int(v) - 38
					} else {
						sum += int(v) - 96
					}

					break out
				}
			}
		}

	}
	//fmt.Printf("%#U", runeArray)
	//part 1 answer
	fmt.Println(sum)
	sumB := 0
	for i := 0; i < len(stringArray); i += 3 {
	outter:
		for _, v := range stringArray[i] {
			for _, b := range stringArray[i+1] {
				if v == b {
					for _, n := range stringArray[i+2] {
						if v == n {
							//fmt.Println(v)
							if v < 96 {
								sumB += int(v) - 38
							} else {
								sumB += int(v) - 96
							}
							break outter
						}
					}
				}
			}
		}

	}
	fmt.Println(sumB)

}
