package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day19() {

	inputText, err := os.Open("input/day19.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer inputText.Close()

	scanner := bufio.NewScanner(inputText)

	robotCosts := make(map[int]robotCost, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)
		blueprint, _ := strconv.Atoi(fields[1])
		ore, _ := strconv.Atoi(fields[6])
		clay, _ := strconv.Atoi(fields[12])
		obs1, _ := strconv.Atoi(fields[18])
		obs2, _ := strconv.Atoi(fields[21])
		geo1, _ := strconv.Atoi(fields[27])
		geo2, _ := strconv.Atoi(fields[30])

		robotCosts[blueprint] = robotCost{ore, clay, obs1, obs2, geo1, geo2}

	}
	ore := 0
	clay := 0
	obsidian := 0
	geode := 0

	oreRobot := 1
	clayRobot := 0
	obsidianRobot := 0
	geodeRobot := 0

	i := 1

	for minute := 1; minute <= 24; minute++ {

		oreRobotBuilding := 0
		clayRobotBuilding := 0
		obsidianRobotBuilding := 0
		geodeRobotBuilding := 0

		if robotCosts[i].geo1 <= ore && robotCosts[i].geo2 <= obsidian {
			ore -= robotCosts[i].geo1
			obsidian -= robotCosts[i].geo2
			geodeRobot++
			geodeRobotBuilding = 1
		}
		if robotCosts[i].obs1 <= ore && robotCosts[i].obs2 <= clay && obsidian+obsidianRobot*2 < robotCosts[i].geo2 {
			ore -= robotCosts[i].obs1
			clay -= robotCosts[i].obs2
			obsidianRobot++
			obsidianRobotBuilding = 1
		}
		if robotCosts[i].clay <= ore && clay+clayRobot*2 < robotCosts[i].obs2 {
			ore -= robotCosts[i].clay
			clayRobot++
			clayRobotBuilding = 1
		}
		if robotCosts[i].ore <= ore {
			ore -= robotCosts[i].ore
			oreRobot++
			oreRobotBuilding = 1
		}

		ore += oreRobot - oreRobotBuilding
		clay += clayRobot - clayRobotBuilding
		obsidian += obsidianRobot - obsidianRobotBuilding
		geode += geodeRobot - geodeRobotBuilding
		fmt.Println("")
		fmt.Println("min", minute)
		fmt.Println("res", ore, clay, obsidian, geode)
		fmt.Println("robots", oreRobot, clayRobot, obsidianRobot, geodeRobot)

	}

}

type robotCost struct {
	ore  int
	clay int
	obs1 int
	obs2 int
	geo1 int
	geo2 int
}
