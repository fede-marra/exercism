// Package weather say the weather condition of same citys.
package weather

// CurrentCondition is useful for save the condition are on this moment.
var CurrentCondition string
//CurrentLocation is useful for save your location.
var CurrentLocation string

// Forecast evaluate through your Current Location and Current Condition which is weather condition now.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
