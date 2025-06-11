package main

import "log"

func main() {
	// Create instances
	currentDisplay := &CurrentCondition{}
	staticsDisplay := &Statics{}
	forecastDisplay := &Forecast{}

	weatherStation := &WeatherStation{
		pressure:    12.3,
		humidity:    0.34,
		temperature: 25.7,
	}
	// Escenario #1 (none subscribed)
	log.Println("********** Escenario #1 **********")
	currentDisplay.Display()
	staticsDisplay.Display()
	forecastDisplay.Display()

	// Escenario #2 (All subscribed)
	log.Println("********** Escenario #2 **********")
	weatherStation.RegisterObserver(currentDisplay)
	weatherStation.RegisterObserver(staticsDisplay)
	weatherStation.RegisterObserver(forecastDisplay)

	// update measurements
	weatherStation.humidity = 0.89
	weatherStation.pressure = 24.2
	weatherStation.temperature = 28.9
	weatherStation.NotifyObservers()

	currentDisplay.Display()
	staticsDisplay.Display()
	forecastDisplay.Display()

	// Escenario #3 (statics removed)
	log.Println("********** Escenario #3 **********")
	weatherStation.RemoveObserver(staticsDisplay)

	weatherStation.humidity = 1.23
	weatherStation.pressure = 34.6
	weatherStation.temperature = 31.2
	weatherStation.NotifyObservers()

	currentDisplay.Display()
	staticsDisplay.Display()
	forecastDisplay.Display()

	// Escenario #4 (forecast removed)
	log.Println("********** Escenario #4 **********")
	weatherStation.RemoveObserver(forecastDisplay)

	weatherStation.humidity = 0.12
	weatherStation.pressure = 12.2
	weatherStation.temperature = 18.9
	weatherStation.NotifyObservers()

	currentDisplay.Display()
	staticsDisplay.Display()
	forecastDisplay.Display()

	// Escenario #5 Nothing changed
	log.Println("********** Escenario #5 **********")
	weatherStation.humidity = 1
	weatherStation.pressure = 2
	weatherStation.temperature = 3
	weatherStation.NotifyObservers()

	currentDisplay.Display()
	staticsDisplay.Display()
	forecastDisplay.Display()
}
