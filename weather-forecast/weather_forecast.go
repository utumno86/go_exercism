// Package weather handles information for forecasting the weather.
package weather

// CurrentCondition stores info about the current wenather conditions at your location.
var CurrentCondition string
// CurrentLocation stores info about the current wenather conditions at your location.
var CurrentLocation string

// Forecast determines the current wheather conditions at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
