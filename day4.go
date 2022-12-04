package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	inputText, err := os.Open("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	overlap := 0
	overlapB := 0

	for scanner.Scan() {

		sectionPerElf := strings.Split(scanner.Text(), ",")
		elfOne := strings.Split(sectionPerElf[0], "-")
		elfTwo := strings.Split(sectionPerElf[1], "-")
		elfOneZero, _ := strconv.Atoi(elfOne[0])
		elfOneOne, _ := strconv.Atoi(elfOne[1])
		elfTwoZero, _ := strconv.Atoi(elfTwo[0])
		elfTwoOne, _ := strconv.Atoi(elfTwo[1])

		if (elfOneZero <= elfTwoZero && elfOneOne >= elfTwoOne) || (elfOneZero >= elfTwoZero && elfOneOne <= elfTwoOne) {
			overlap++
			// fmt.Println(elfOne[0], elfOne[1], elfTwo[0], elfTwo[1])

		}

		if elfOneZero >= elfTwoZero && elfOneZero <= elfTwoOne || elfOneOne >= elfTwoZero && elfOneOne <= elfTwoOne || elfTwoZero >= elfOneZero && elfTwoZero <= elfOneOne || elfTwoOne >= elfOneZero && elfTwoOne <= elfOneOne {
			overlapB++
			fmt.Println(elfOne[0], elfOne[1], elfTwo[0], elfTwo[1])

		}
	}
	fmt.Println(overlap)
	fmt.Println(overlapB)

}
