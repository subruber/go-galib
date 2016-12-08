/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Example of uing genome_ordered_int.

Find the genome with max sum.
*/
package main

import (
	"fmt"
	"github.com/thoj/go-galib"
	"math/rand"
	"os"
	"time"
)

var calls int

// Boring fitness/score function.
func score(g *ga.GAIntOfAlphabetGenome) float64 {
	var total int
	for _, v := range g.Gene {
		total += 4 - v
	}
	calls++
	return float64(total)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	param := ga.GAParameter{
		Initializer: new(ga.GARandomInitializer),
		Selector:    ga.NewGATournamentSelector(0.7, 5),
		Breeder:     new(ga.GA2PointBreeder),
		Mutator:     new(ga.GARandomMutator),
		PMutate:     0.1,
		PBreed:      0.7}

	gao := ga.NewGA(param)

	alphabet := []int{1, 2, 3, 4}
	weight := []int{1, 2, 2, 4}
	genome, err := ga.NewIntOfAlphabetGenome(alphabet, weight, []int{1, 2, 3, 4, 4, 2, 2, 1}, score)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gao.Init(10, genome) //Total population
	gao.OptimizeUntil(func(best ga.GAGenome) bool {
		return best.Score() <= 5
	})
	gao.PrintTop(5)

	best := gao.Best().(*ga.GAIntOfAlphabetGenome)
	fmt.Printf("%s = %f\n", best, best.Score())
	fmt.Printf("Calls to get the best score = %d\n", calls)
}
