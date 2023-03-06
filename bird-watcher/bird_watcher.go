package birdwatcher

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var suma int
	for i := 0; i < len(birdsPerDay) ; i++ {
		fmt.Println(birdsPerDay[i])
		suma += birdsPerDay[i]
	}
	return suma
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int
	fmt.Println(week)
	switch week {
	
	case  1:
		for i := 0; i <= 6; i ++ {
			sum += birdsPerDay[i]}
	case 2:
		for i := 7; i <= 14; i ++ {
			sum += birdsPerDay[i]}
	case 3:
		for i := 15; i <= 22; i ++ {
			sum += birdsPerDay[i]}
	default:
		for i := 23; i <= len(birdsPerDay); i ++ {
			sum += birdsPerDay[i]}
	}
	return sum	
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay);i = i + 2 {
		birdsPerDay[i] +=1
	}
	return birdsPerDay
}
