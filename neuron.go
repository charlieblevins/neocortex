package main

import (
	"fmt"
)

type Neuron struct {
	dendrite Dendrite
	axon     Axon
}

func (n *Neuron) init() {
	fmt.Println("I'm a neuron")

	// TODO: listen for spike on dendrite and
	// pass on to axon
	n.dendrite.register(n.axon)
}

func makeNeuron() Neuron {
	return Neuron{Dendrite{}, Axon{}}
}
