package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("puzzle.txt")
	handleErr(err)

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var leftLocations []int
	var rightLocations []int

	for scanner.Scan() {
		locations := strings.Split(scanner.Text(), "   ")

		leftLocations = readLocations(leftLocations, locations[0])
		rightLocations = readLocations(rightLocations, locations[1])
	}

	sort.Ints(leftLocations)
	sort.Ints(rightLocations)

	var distanceTotal int

	for index, locIdLeft := range leftLocations {

		distance := rightLocations[index] - locIdLeft
		if distance < 0 {
			distance = distance * -1
		}
		distanceTotal += distance
	}

	fmt.Printf("Total distance: %d\n", distanceTotal)
}

func readLocations(locations []int, locCode string) []int {
	locId, err := strconv.Atoi(locCode)

	handleErr(err)

	locations = append(locations, locId)
	return locations
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
