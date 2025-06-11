package main

import "log"

// Observer is who receive the information
type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}

// Display is the interface for the Displays to show information
type Display interface {
	Display()
}

// Verbose Implement Observer Interface
var _ Observer = &CurrentCondition{}
var _ Observer = &Statics{}
var _ Observer = &Forecast{}

// Verbose Implement Display Interface
var _ Display = &CurrentCondition{}
var _ Display = &Statics{}
var _ Display = &Forecast{}

// CurrentCondition implements Observer because is notify when the state changes
type CurrentCondition struct {
	pressure    float32
	humidity    float32
	temperature float32
}

func (c *CurrentCondition) Update(temp float32, humidity float32, pressure float32) {
	c.humidity = humidity
	c.pressure = pressure
	c.temperature = temp
}

func (c CurrentCondition) Display() {
	log.Println("Displaying Current conditions:")
	log.Printf("Humidity measurement: %.2f \n", c.humidity)
	log.Printf("Preassure measurement: %.2f \n", c.pressure)
	log.Printf("Temperature measurement: %.2f \n", c.temperature)
	log.Println(" ")
}

// Statics implements Observer because is notify when the state changes
type Statics struct {
	humidity float32
}

func (s *Statics) Update(temp float32, humidity float32, pressure float32) {
	s.humidity = humidity
}

func (s Statics) Display() {
	log.Println("Displaying Statics conditions:")
	log.Printf("Humidity measurement: %.2f \n", s.humidity)
	log.Println(" ")
}

// Forecast implements Observer because is notify when the state changes
type Forecast struct {
	pressure float32
}

func (f *Forecast) Update(temp float32, humidity float32, pressure float32) {
	f.pressure = pressure
}

func (f Forecast) Display() {
	log.Println("Displaying Forecast conditions:")
	log.Printf("Preassure measurement: %.2f \n", f.pressure)
	log.Println(" ")
}
