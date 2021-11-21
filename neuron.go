package main

import (
	"fmt"
)

type Neuron struct {
	dendrite Dendrite
	axon     Axon
}

func (n *Neuron) spike() {
	spike := "spike!"
	fmt.Println("Nueron: " + spike)
	n.axon.send(spike)
}

func makeNeuron(a Axon, d Dendrite) Neuron {
	return Neuron{d, a}
}
