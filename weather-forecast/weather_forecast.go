// Package weather includes the following functions:
// - Forecast - returns the current weather condition for the given city.
package weather

// CurrentCondition: The CurrentLocation's current weather condition.
var CurrentCondition string
// CurrentLocation: The current city where the weather is being forecasted for.
var CurrentLocation string

// Forecast returns the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
