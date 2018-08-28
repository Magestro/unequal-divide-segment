package main

import (
	"log"
	"math/rand"
	"time"
)

func UnequalDivideSegment(count uint, length, min, max float64) []float64 {
	log.Printf("count: %d; length: %f; min: %f; max: %f", count, length, min, max)
	a := randFloats(min, max, int(count))
	s := sum(a...)
	log.Printf("Sum 1: %f \n", s)

	for i := range a {
		a[i] /= s
		a[i] *= length
		//a[i] *= (max + min) / 2
	}

	log.Printf("Sum 2: %f", sum(a...))

	return a
}

func randFloats(min, max float64, n int) []float64 {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	res := make([]float64, n)
	for i := range res {
		res[i] = min + r.Float64()*(max-min)
	}
	return res
}

func sum(input ...float64) float64 {
	sum := 0.0

	for i := range input {
		sum += input[i]
	}

	return sum
}
