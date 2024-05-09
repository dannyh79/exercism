package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var successPercentage = successRate / 100
	return float64(productionRate) * successPercentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	const minutesPerHour = 60
	var successPercentage = successRate / 100

	return int(float64(productionRate) / minutesPerHour * successPercentage)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const (
		groupSize      = 10
		groupCost      = 95_000
		individualCost = 10_000
	)

	groups := carsCount / groupSize
	individuals := carsCount % groupSize
	return uint(groups*groupCost + individuals*individualCost)
}
