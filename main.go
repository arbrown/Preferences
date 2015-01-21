// Preferences project main.go
package main

import (
	"errors"
	"fmt"
	"sort"
)

// Stores preferences for each Role for a given player
// Trapper, Medic, Support, Assault, & Monster
// 1 is the most preferred, 5 is the least preferred
type RolePreference struct {
	Trapper, Medic, Support, Assault, Monster int
	Name                                      string
}

type RoleAssignment struct {
	Name, Role string
}

// Assign roles to a group of players
func AssignRoles(prefs []RolePreference) ([]RoleAssignment, int, error) {
	if len(prefs) != 5 {
		return nil, 25, errors.New("Need 5 players to assign roles")
	}

	// We have 5 players, refer to them by their slice index
	// p1 = prefs[0] etc...

	bestWeight := 25 // worst case, everyone gets least favorite
	bestSolutions := make([][]int, 1)

	for current := range GeneratePermutations([]int{1, 2, 3, 4, 5}) {
		currentWeight, _ := GetWeight(current, prefs)
		if currentWeight <= bestWeight {
			if currentWeight < bestWeight {
				bestWeight = currentWeight
				bestSolutions = nil
			}
			sliceCopy := make([]int, len(current))
			copy(sliceCopy, current)
			bestSolutions = append(bestSolutions, sliceCopy)
		}
	}

	// for now, just take the first, best solution
	if len(bestSolutions) == 0 {
		return nil, 25, errors.New("Could not find any valid solutions")
	}

	bestSolution := bestSolutions[0]
	assignments := make([]RoleAssignment, len(prefs))

	for i, p := range prefs {
		name, _ := GetRoleName(bestSolution[i])
		assignments[i] = RoleAssignment{Name: p.Name, Role: name}
	}

	return assignments, bestWeight, nil
}

// Get the name associated with a role ID
func GetRoleName(i int) (string, error) {
	// 1 - Trapper
	// 2 - Medic
	// 3 - Support
	// 4 - Assault
	// 5 - Monster
	switch i {
	case 1:
		return "Trapper", nil
	case 2:
		return "Medic", nil
	case 3:
		return "Support", nil
	case 4:
		return "Assault", nil
	case 5:
		return "Monster", nil
	}

	return "", errors.New("Role not found")
}

// Get the preference weight for this role permutation
func GetWeight(current []int, prefs []RolePreference) (int, error) {
	sum := 0
	for i := 0; i < 5; i++ {
		if w, e := GetRoleWeight(current[i], prefs[i]); e == nil {
			sum += w
		} else {
			return 25, e
		}
	}
	return sum, nil
}

// Get an individual player's preference for his given role
func GetRoleWeight(role int, playerPreferences RolePreference) (int, error) {
	// 1 - Trapper
	// 2 - Medic
	// 3 - Support
	// 4 - Assault
	// 5 - Monster
	switch role {
	case 1:
		return playerPreferences.Trapper, nil
	case 2:
		return playerPreferences.Medic, nil
	case 3:
		return playerPreferences.Support, nil
	case 4:
		return playerPreferences.Assault, nil
	case 5:
		return playerPreferences.Monster, nil
	}
	return -1, errors.New("Role not found")
}

// Return every permutation of integers in a slice
func GeneratePermutations(values []int) <-chan []int {
	c := make(chan []int)

	// Make sure our slice of integers is sorted
	sort.Ints(values)

	go func(c chan []int) {
		defer close(c)

		done := false
		for !done {
			// Return the next lexicographical permutation
			retSlice := make([]int, len(values))
			copy(retSlice, values)
			c <- retSlice

			// Find the rightmost value smaller than the one after it
			i := len(values) - 2
			for ; i >= 0; i-- {
				if values[i] < values[i+1] {
					break
				}
			}

			// If we managed to get to -1, values are reverse sorted and we
			// have returned all permutations
			if i == -1 {
				done = true
			} else {
				// find the smallest number greater than values[i] but still
				// after it in the slice (ceil)
				j, ceil := i+1, i+1
				for ; j < len(values); j++ {
					if values[j] > values[i] && values[j] < values[ceil] {
						ceil = j
					}
				}
				// and swap them
				values[i], values[ceil] = values[ceil], values[i]

				// then sort the new subslice
				sort.Ints(values[i+1 : len(values)])

			}
		}
	}(c)

	return c
}

func main() {
	sample := []RolePreference{
		RolePreference{1, 2, 3, 4, 5, "Vincent"},
		RolePreference{5, 4, 3, 2, 1, "Edgar"},
		RolePreference{1, 5, 2, 4, 3, "Lenn0s"},
		RolePreference{1, 2, 3, 4, 5, "Perlmonger42"},
		RolePreference{1, 4, 2, 3, 5, "Walrus"},
	}

	result, weight, _ := AssignRoles(sample)

	fmt.Println(result)
	fmt.Printf("Weight: %d\n", weight)
}
