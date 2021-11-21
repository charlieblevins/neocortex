package main

import "fmt"

type Axon struct {
	terminals []AxonTerminal
}

func (a Axon) send(spike string) {
	fmt.Println("Axon sending " + spike)

	for _, t := range a.terminals {
		t.releaseChemical()
	}
}

func (a Axon) getId() string {
	return "id"
}

func makeAxon(ats []AxonTerminal) Axon {
	return Axon{terminals: ats}
}
