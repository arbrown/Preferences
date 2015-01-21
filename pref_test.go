package main

import "testing"

var prefTests = []struct {
	in     []RolePreference
	out    []RoleAssignment
	weight int
}{
	{
		[]RolePreference{
			{1, 2, 3, 4, 5, "Vincent"},
			{5, 1, 2, 3, 4, "Edgar"},
			{4, 5, 1, 2, 3, "Lenn0s"},
			{3, 4, 5, 1, 2, "Perlmonger42"},
			{2, 3, 4, 5, 1, "Walrus"},
		},
		[]RoleAssignment{
			{"Vincent", "Trapper"},
			{"Edgar", "Medic"},
			{"Lenn0s", "Support"},
			{"Perlmonger42", "Assault"},
			{"Walrus", "Monster"},
		},
		5,
	},
}

func TestPrefs(t *testing.T) {
	for _, pt := range prefTests {
		actual, _, _ := AssignRoles(pt.in)
		if !resultsEqual(actual, pt.out) {
			t.Error("Error")
		}
	}
}

func benchmarkPermutation(s []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum := 0
		for p := range GeneratePermutations(s) {
			sum += p[0]
		}
	}
}

func Benchmark10(b *testing.B) {
	benchmarkPermutation([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, b)
}

func Benchmark5(b *testing.B) {
	benchmarkPermutation([]int{0, 1, 2, 3, 4}, b)
}

func Benchmark3(b *testing.B) {
	benchmarkPermutation([]int{0, 1, 2}, b)
}

func resultsEqual(actual, expected []RoleAssignment) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		a, e := actual[i], expected[i]
		if a.Name != e.Name || a.Role != e.Role {
			return false
		}
	}
	return true
}
