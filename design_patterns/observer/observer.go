package main

import (
	"fmt"
)

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getID() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func (item *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", item.name)
	item.available = true
	item.broadcast()
}

func (item *Item) register(observer Observer) {
	item.observers = append(item.observers, observer)
}

func (item *Item) broadcast() {
	for _, observer := range item.observers {
		observer.updateValue(item.name)
	}
}

type Email struct {
	id string
}

func (email *Email) getID() string {
	return email.id
}

func (email *Email) updateValue(value string) {
	fmt.Printf("Sending Email -- %s available from Client %s\n", value, email.id)
}

func main() {
	item := &Item{name: "Nvidia RTX 4080"}
	firstObserver := &Email{id: "25aa"}
	secondObserver := &Email{id: "32bb"}
	item.register(firstObserver)
	item.register(secondObserver)
	item.UpdateAvailable()
}
