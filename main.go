package main

import (
	"log"
	"os"

	"github.com/Isnor/beesolver/beesolver"
	"github.com/urfave/cli"
)

func main() {
	var wordListPath, requiredLetter string

	solver := &beesolver.BeeSolver{}
	// parse args
	cli := &cli.App{
		Name: "beesolver",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "l",
				Usage:    "The 6 letters of the puzzle",
				Required: true,
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
			var err error
			solver, err = beesolver.NewBeeSolver(ctx.StringSlice("l"), wordListPath, requiredLetter, 4, 16)
			return err
		},
	}

	// parse the arguments to build the solver
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// run solver
	solutions, err := solver.Solve()
	log.Printf("%s\n%v", solutions, err)
}
