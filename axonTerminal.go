package main

import (
	"fmt"
	"time"
)

const CHEMICAL_RELEASE_DELAY_MS = 500

type AxonTerminal struct {
	// TODO: is this optional? does every axon termina connect to anothert dendrite??
	externalDendrite Dendrite
}

func (at *AxonTerminal) releaseChemical() {
	fmt.Println("Axon Terminal: begin release chemical")

	// this process takes "milliseconds" to complete
	// TODO: find out what the normal amount of time delay is,
	// TODO: find out if time delay is variable and what causes it
	// to change
	time.Sleep(CHEMICAL_RELEASE_DELAY_MS * time.Millisecond)

	fmt.Println("Axon Terminal: finish release chemical")
}

func makeAxonTerminal(d Dendrite) AxonTerminal {
	return AxonTerminal{externalDendrite: d}
}
