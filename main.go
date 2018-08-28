package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var count uint
var length float64
var min float64
var max float64

func init() {
	flag.UintVar(&count, "count", 24, "total count of segments")

	flag.Float64Var(&length, "length", 24, "whole length, must be positive")

	flag.Float64Var(&min, "min", 1, "min segment length, must be zero or positive")

	flag.Float64Var(&max, "max", 5, "max segment length, must be greater than min ")

	flag.Parse()
}

func main() {
	validateFlags()

	segments := UnequalDivideSegment(count, length, min, max)

	text := []string{}
	for i := range segments {
		number := segments[i]
		t := strconv.FormatFloat(number, 'f', 6, 64)
		text = append(text, t)
	}

	// Join our string slice.
	result := strings.Join(text, " ")
	fmt.Println(result)
}

func validateFlags() {
	if length <= 0 {
		panic(fmt.Errorf("arg `length` must be positive"))
	}

	if min < 0 {
		panic(fmt.Errorf("arg `min` must be zero or positive"))
	}

	if min > max {
		panic(fmt.Errorf("arg `max` must be greater than min"))
	}

	if length/float64(count) < min {
		panic(fmt.Errorf("arg `min` must be lower than `lenght / count` (%f)", length/float64(count)))
	}
}
