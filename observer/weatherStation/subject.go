package main

import (
	"log"
)

// Subject is who publish the information
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// WeatherStation implements Subject because has the state
type WeatherStation struct {
	pressure    float32
	humidity    float32
	temperature float32

	// List of observers receiving data
	observers []Observer
}

// RegisterObserver adds observer to the observer list
func (w *WeatherStation) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

// RemoveObserver removes observer from the observer list
func (w *WeatherStation) RemoveObserver(observer Observer) {
	// found the index of the observer to be deleted
	observerIndex := -1
	for i, obs := range w.observers {
		if obs == observer {
			observerIndex = i
			break
		}
	}

	// Remove the observer if was found
	if observerIndex > 0 {
		w.observers = append(w.observers[:observerIndex], w.observers[observerIndex+1:]...)
		log.Printf("Observer: %d was deleted\n", observerIndex)
		return
	}
}

// NotifyObservers execute the update method of its observer
func (w WeatherStation) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}
