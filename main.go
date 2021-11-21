package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to your new neocortex")

	d2 := makeDendrite()
	makeNeuron(makeAxon([]AxonTerminal{}), d2)

	d1 := makeDendrite()
	n1 := makeNeuron(
		makeAxon([]AxonTerminal{makeAxonTerminal(d2)}),
		d1,
	)

	n1.spike()
}
