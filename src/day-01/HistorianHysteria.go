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

	input, err := os.Open("src/day-01/puzzle.txt")
	handleErr(err)

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var locationIds []int
	var chiefNotes []int

	for scanner.Scan() {
		locations := strings.Split(scanner.Text(), "   ")

		locationIds = readLocations(locationIds, locations[0])
		chiefNotes = readLocations(chiefNotes, locations[1])
	}

	sort.Ints(locationIds)
	sort.Ints(chiefNotes)

	// Day One
	calculateDistances(locationIds, chiefNotes)
	calculateSimilarityScore(chiefNotes, locationIds)
}

func calculateSimilarityScore(chiefNotes []int, locationIds []int) {
	locIdCount := make(map[int]int)

	for _, noteLocId := range chiefNotes {
		locIdCount[noteLocId] += 1
	}

	var distance int
	for _, locId := range locationIds {
		distance += locId * locIdCount[locId]
	}
	fmt.Printf("Total similarity score %d", distance)
}

func calculateDistances(locationIds []int, chiefNotes []int) {
	var distanceTotal int

	for index, locIdLeft := range locationIds {

		distance := chiefNotes[index] - locIdLeft
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
