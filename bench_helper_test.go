package mapset

import (
	"fmt"
	"math/rand"
)

func nrand(n int) []int {
	is := make([]int, n)
	for i := range is {
		is[i] = rand.Int()
	}
	return is
}

func buildRandomSet(n int, safe bool) Set[int] {
	if safe {
		return NewSet(nrand(n)...)
	} else {
		return NewThreadUnsafeSet(nrand(n)...)
	}
}

func buildCaseName(n int, safe bool) string {
	var settype string
	if safe {
		settype = "Safe"
	} else {
		settype = "Unsafe"
	}

	return fmt.Sprintf("%d_%s", n, settype)
}

type benchCase struct {
	n    int
	safe bool
}

func buildBenchCases(ns ...int) []benchCase {
	cases := []benchCase{}
	for _, n := range ns {
		cases = append(cases, benchCase{n, true})
		cases = append(cases, benchCase{n, false})
	}
	return cases
}
