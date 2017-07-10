package main

import (
	"fmt"
)

func main() {
	set := []float64{3, 5, 5, 5, 5, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	ocurrenceMap := make(map[float64]int, 0)
	for _, val := range set {
		if _, ok := ocurrenceMap[val]; ok {
			ocurrenceMap[val]++
		} else {
			ocurrenceMap[val]++
		}
	}

	var modeKeys, auxModeKeys []float64
	var jumpBar int
	for key, val := range ocurrenceMap {
		if val > jumpBar {
			jumpBar = val
			modeKeys = append(auxModeKeys, key)
		} else if val == jumpBar {
			modeKeys = append(modeKeys, key)
		}

	}

	mode := ocurrenceMap[modeKeys[0]]

	fmt.Println("Mode Key(s):", modeKeys, "Mode:", mode)

	// fmt.Println("Mode: ", mode)
}
