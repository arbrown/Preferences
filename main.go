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
func AssignRoles(prefs []RolePreference) ([]RoleAssignment, error) {
	if len(prefs) != 5 {
		return nil, errors.New("Need 5 players to assign roles")
	}

	// We have 5 players, refer to them by their slice index
	// p1 = prefs[0] etc...
	return nil, nil
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
			fmt.Println("About to return ", values)
			retSlice := make([]int, len(values))
			copy(retSlice, values)
			c <- values
			fmt.Println("Just returned ", retSlice)

			//fmt.Println("values = ", values)

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
				values[i] = values[i] ^ values[ceil]
				values[ceil] = values[ceil] ^ values[i]
				values[i] = values[i] ^ values[ceil]

				// then sort the new subslice
				sort.Ints(values[i+1 : ceil+1])

			}
		}
	}(c)

	return c
}

func main() {
	i := 1
	for combo := range GeneratePermutations([]int{1, 2, 3, 4, 5}) {
		fmt.Println(i, " - ", combo)
		i++
	}
	//c := GeneratePermutations([]int{1, 2, 3, 4, 5})
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
}
