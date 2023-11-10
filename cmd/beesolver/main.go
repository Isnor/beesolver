package main

import (
	"log"
	"os"
	"strings"

	"github.com/Isnor/beesolver"
	"github.com/urfave/cli"
)

func main() {
	var wordListPath, letters, requiredLetter string

	solver := &beesolver.BeeSolver{}
	// parse args
	cli := &cli.App{
		Name: "beesolver",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "letters",
				Usage:       "The 6 letters of the puzzle, comma separated",
				Destination: &letters,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "middle",
				Usage:       "the middle, required letter",
				Required:    true,
				Destination: &requiredLetter,
			},
			&cli.StringFlag{
				Name:        "dict",
				Usage:       "path to the words list",
				Value:       "/usr/share/dict/words",
				Destination: &wordListPath,
			},
		},
		Action: func(ctx *cli.Context) error {
			allowedLetters := strings.Split(letters, ",")
			var err error
			solver, err = beesolver.NewBeeSolver(allowedLetters, wordListPath, requiredLetter, 4, 16)
			return err
		},
	}

	// parse the arguments to build the solver
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// run solver
	solutions, _ := solver.Solve()
	if len(solutions) == 0 {
		log.Fatal("no solutions found")
	}
	log.Printf("found solutions: %s\n", solutions)
}
