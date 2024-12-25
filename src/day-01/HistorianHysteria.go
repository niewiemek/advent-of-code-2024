package main

import (
	"../helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var locationIds []int
	var chiefNotes []int

	var fileHandler helper.LineHandler = func(line string) {
		locations := strings.Split(line, "   ")

		locationIds = readLocations(locationIds, locations[0])
		chiefNotes = readLocations(chiefNotes, locations[1])
	}

	helper.ReadFile("day-01/puzzle.txt", fileHandler)

	sort.Ints(locationIds)
	sort.Ints(chiefNotes)

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

	helper.HandleErr(err)

	locations = append(locations, locId)
	return locations
}
