package beesolver_test

import (
	"sort"
	"strings"
	"testing"

	"isnor.ca/beesolver/beesolver"
)

func TestBee(t *testing.T) {
	solver, err := beesolver.NewBeeSolver(
		[]string{"d", "r", "m", "i", "t", "y"},
		"/usr/share/dict/words",
		"o",
		4,
		16,
	)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(solver.ValidWordRegex.String())

	solutions, err := solver.Solve()
	t.Logf("%s\n%v", solutions, err)

	t.Fail()
}

func TestBee2(t *testing.T) {
	solver, err := beesolver.NewBeeSolver(
		[]string{"d", "r", "m", "i", "t", "y"},
		"../words_alpha.txt",
		"o",
		4,
		16,
	)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(solver.ValidWordRegex.String())

	solutions, err := solver.Solve()
	sort.Strings(solutions)
	strings.Join(solutions, "\n")
	t.Logf("\n%s\n%v\n", strings.Join(solutions, "\n"), err)

	t.Fail()
}


func TestBee3(t *testing.T) {
	solver, err := beesolver.NewBeeSolver(
		[]string{"m", "c", "i", "k", "a", "n"},
		"../words_alpha.txt",
		"g",
		7,
		17,
	)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(solver.ValidWordRegex.String())

	solutions, err := solver.Solve()
	sort.Strings(solutions)
	strings.Join(solutions, "\n")
	t.Logf("\n%s\n%v\n", strings.Join(solutions, "\n"), err)

	t.Fail()
}