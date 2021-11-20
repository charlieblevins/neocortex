package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to your new neocortex")

	n := makeNeuron()
	n.init()
	n.dendrite.spike()
}
