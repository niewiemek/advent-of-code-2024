package main

import (
	"../helper"
	"fmt"
	"slices"
	"strings"
)

const (
	DamperEnabled = true
)

func main() {

	var safeReports int
	var handler helper.LineHandler = func(line string) {
		var report = helper.ToIntArray(strings.Split(line, " "))
		if isSafe(report) {
			safeReports++
		} else {
			if DamperEnabled {
				for i := 0; i < len(report); i++ {
					clonedReport := helper.Clone(report)
					withLevelRemoved := slices.Delete(clonedReport, i, i+1)
					if isSafe(withLevelRemoved) {
						safeReports++
						break
					}
				}
			}
		}
	}

	helper.ReadFile("day-02/puzzle.txt", handler)
	fmt.Printf("\nSafe reports count: %d\n", safeReports)
}

func isSafe(report []int) bool {
	fmt.Print("\n-+-+-\n\n")
	fmt.Printf("Analyzing %v\n", report)
	var (
		directionVector    helper.Direction
		levelSafetyResults = make([]bool, len(report)-1)
		prevLevelValue     = report[0]
	)
	for level, value := range report {
		if level > 0 {
			fmt.Printf("[%d, %d] | ", prevLevelValue, value)
			levelDiff, direction := helper.AbsDiff(prevLevelValue, value)
			if level == 1 {
				directionVector = direction
			}
			directionVector *= directionVector
			fmt.Printf("Direction: %v | ", direction)
			fmt.Printf("Direction vector: %v | ", directionVector)
			constantDirection := directionVector == direction
			fmt.Printf("Constant direction: %v | ", constantDirection)
			levelDiffIsSafe := levelDiff >= 1 && levelDiff <= 3
			fmt.Printf("Safe level difference: %v\n", levelDiffIsSafe)

			prevLevelValue = value
			levelSafetyResults[level-1] = constantDirection && levelDiffIsSafe
		}
	}

	isSafeReport := helper.Reduce(levelSafetyResults, func(acc bool, currentVal bool) bool {
		return acc && currentVal
	}, true)

	fmt.Printf("\n%v, isSafe %v\n", report, isSafeReport)
	return isSafeReport
}
