// Package weather: Good package.
package weather

// CurrentCondition: Good condition.
var CurrentCondition string

// CurrentLocation: Good location.
var CurrentLocation string

// Forecast: Good function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
