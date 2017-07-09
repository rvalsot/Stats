package main

import (
	"fmt"
)

func main() {

	ocurrenceMap := make(map[float64]int, 0)

	set := []float64{3, 3, 3, 5, 5, 5, 5, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	for _, val := range set {
		if _, ok := ocurrenceMap[val]; ok {
			ocurrenceMap[val]++
		} else {
			ocurrenceMap[val]++
		}
	}

	var modeKeys []float64
	var jumpBar int

	for key, val := range ocurrenceMap {
		fmt.Println(key, ":", val)
		if val >= jumpBar {
			var auxModeKeys []float64
			jumpBar = val

			fmt.Println(auxModeKeys, "jumpBar →", val)
			modeKeys = append(modeKeys, key)
		}
	}

	fmt.Println("Final:", modeKeys)

	for key, val := range ocurrenceMap {
		fmt.Println(key, ":", val)
	}

	// fmt.Println("Mode: ", mode)
}
