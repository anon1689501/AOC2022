package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day16() {

	inputText, err := os.Open("input/day16.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	valveList := make([]valve, 0)

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.ReplaceAll(line, "rate=", "")
		line = strings.ReplaceAll(line, ";", "")
		line = strings.ReplaceAll(line, ",", "")
		//fmt.Println(line)
		fields := strings.Fields(line)
		name := fields[1]
		rate, _ := strconv.Atoi(fields[4])
		leadsTo := make([]string, 0)
		for i := 9; i < len(fields); i++ {
			leadsTo = append(leadsTo, fields[i])
		}
		valveList = append(valveList, valve{name, rate, leadsTo, false})
		//fmt.Println(valveList)

	}

	flowStart("AA", 1, valveList)

}

type valve struct {
	name    string
	rate    int
	leadsTo []string
	open    bool
}

func flowStart(name string, step int, valveList []valve) int {

	openVal := 0
	flowing := 0
	travel := make([]int, 0)

	for _, v := range valveList {
		if v.open {
			flowing += v.rate //sum currently open valves
		}
	}
	if step >= 30 {
		return flowing
	}

	for _, v := range valveList { //open current valve
		if v.name == name && !v.open {
			openVal = flowStart(name, step+1, valveList) //move on
		}
	}

	for _, v := range valveList {
		if v.name == name {
			for _, v2 := range v.leadsTo {
				travel = append(travel, flowStart(v2, step+1, valveList))

			}
		}

	}

	//move to list of valves at the end
	for _, v := range travel {
		if v > openVal {
			openVal = v
		}
	}
	return flowing + openVal

}
