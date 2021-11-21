package main

import "fmt"

type Dendrite struct {
	observerList []observer
}

// TODO: needs voltage param??
func (d *Dendrite) spike() {
	fmt.Println("Dendrite: action potential received")

	// TODO: does this need a time delay or does the axon? or
	// should the neuron create the delay?
	d.notifyAll()
}

func (d *Dendrite) register(o observer) {
	d.observerList = append(d.observerList, o)
}

func (d *Dendrite) notifyAll() {
	for _, observer := range d.observerList {
		observer.update("spike!")
	}
}

func makeDendrite() Dendrite {
	return Dendrite{}
}
