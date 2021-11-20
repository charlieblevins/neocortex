package main

import "fmt"

type Axon struct{}

func (a Axon) spike() {
	fmt.Println("Axon spiked")
}

func (a Axon) update(message string) {
	fmt.Println("Axon received: " + message + " from dendrite")
	// TODO: this should not always happen - what factors are at play?
	a.spike()
}

func (a Axon) getId() string {
	return "id"
}
