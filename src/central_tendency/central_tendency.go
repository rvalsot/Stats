package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
)

// Structs ____________________________________________________________________

// DiscreteSet64 is a struct, with a name for a slice of int64 data and attached methods for statistics analysis.
type DiscreteSet64 struct {
	Name string
	Data []int64
}

// ContinuousSet64 is a struct, with a name for a slice of float64 data and attached methods for statistics analysis.
type ContinuousSet64 struct {
	Name string
	Data []float64
}

// Methods ____________________________________________________________________

func (set ContinuousSet64) emptyDataSet() (nan float64, err error) {
	if len(set.Data) == 0 {
		nan = math.NaN()
		err = errors.New("Empty data set, could not perform operations")
	}
	return nan, err
}

// Sort takes a ContinuousSet64 struct, sorts in ascending order its data and returns it for further usage
func (set ContinuousSet64) Sort() ContinuousSet64 {
	sort.Float64s(set.Data)

	var orderedSet = ContinuousSet64{set.Name, set.Data}
	return orderedSet
}

// Median method takes a ContinuousSet64 struct, returns its Median value or err if empty
func (set ContinuousSet64) Median() (median float64, err error) {
	// Error if empty set, returns Not-a-number
	set.emptyDataSet()

	// Median cases: even or odd number of elements
	length := len(set.Data)
	if length%2 == 0 {
		median = set.Data[length/2]*0.5 + set.Data[length/2-1]*0.5
	} else {
		position := int64(length/2 + 1)
		median = set.Data[position]
	}

	err = nil
	return median, err
}

// Mean returns the arithmetic average of a ContinuousSet64 or a Not-a-number if there is no data.
func (set ContinuousSet64) Mean() (mean float64, err error) {
	if len(set.Data) == 0 {
		mean = math.NaN()
		err = errors.New("Empty data set")
		return mean, err
	}

	var sum float64
	length := len(set.Data)
	for _, v := range set.Data {
		sum = sum + v
	}

	mean = sum / float64(length)
	err = nil
	return mean, err
}

// Variance returns the variance of a ContinuousSet64 or a Not-a-number if there is no data
func (set ContinuousSet64) Variance() (variance float64, err error) {
	return 1, nil
}

// Main _______________________________________________________________________

func main() {

	aSlice := ContinuousSet64{"vaca", []float64{1, 2, 3, 4, 4, 4, 12, 6, 5, 3, 3, 4, 48, 0, 9, 2, 2, 12, 43, 89, 45, 6, 7, 8, 9, 4, 6, 7, 4, 4, 4, 43, 3, 3, 23, 1, 23, 32, 3, 8, 9, 3, 201, 23, 4, 7, 8, 9, 1, 2, 3, 4, 5, 6}}

	fmt.Println(len(aSlice.Data))
	fmt.Println("Ascendant order: ", aSlice.Sort())

	aMedian, err := aSlice.Median()
	if err != nil {
		log.Println(err)
	}

	aMean, err := aSlice.Mean()
	if err != nil {
		log.Println(err)
	}

	emptySlice := ContinuousSet64{"Empty", []float64{}}

	emptyMedian, err := emptySlice.Median()
	if err != nil {
		log.Println(emptyMedian)
		log.Println(err)
	}

	fmt.Println("Median: ", aMedian)
	fmt.Println("Mean:", aMean)

}
