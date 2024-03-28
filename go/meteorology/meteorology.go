package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (u TemperatureUnit) String() string {
	if u == 0 {
		return "°C"
	} else {
		return "°F"
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (u SpeedUnit) String() string {
	if u == 0 {
		return "km/h"
	} else {
		return "mph"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (t MeteorologyData) String() string {
	s := "%s: %v, Wind %s at %v, %d%% Humidity"
	return fmt.Sprintf(s, t.location, t.temperature, t.windDirection, t.windSpeed, t.humidity)
}
