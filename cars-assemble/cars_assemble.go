package cars

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return  (successRate * float64(productionRate)) / 100
	
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	totalcar := ((successRate * float64(productionRate)) / 100)
	fmt.Println(totalcar)
	return int(totalcar / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tenblock := carsCount / 10
	cost1 := (tenblock * 95000)
	individual := carsCount - (tenblock * 10)
	cost2 := individual * 10000
	return uint(cost1 + cost2)
}
