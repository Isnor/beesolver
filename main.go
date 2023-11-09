package main

import (
	"log"

	"isnor.ca/beesolver/beesolver"
)

func main() {
	// parse args
	// create solver
	solver, err := beesolver.NewBeeSolver(
		[]string{"d", "r", "m", "i", "t", "y"},
		"/usr/share/dict/words",
		"o",
		4,
		16,
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	// run solver

	solutions, err := solver.Solve()
	log.Printf("%s\n%v", solutions, err)
}
