package beesolver_test

import (
	"testing"

	"github.com/Isnor/beesolver/beesolver"
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

	if len(solutions) <= 0 {
		t.Error("there should have been solutions found")
	}
}
