package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var c = 0
	for _, b := range birdsPerDay {
		c += b
	}
	return c
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	const (
		daysPerWeek = 7
		dateOffset  = 1
	)

	s := (week - dateOffset) * daysPerWeek
	e := s + daysPerWeek - dateOffset
	var c = 0
	for i := s; i <= e; i++ {
		c += birdsPerDay[i]
	}
	return c
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	const countToBeFixed = 1
	var isDateMistakeMade = func(d int) bool {
		var everySecondDay = 2
		return d%everySecondDay == 0
	}

	for i := range birdsPerDay {
		if isDateMistakeMade(i) {
			birdsPerDay[i] += countToBeFixed
		}
	}
	return birdsPerDay
}
