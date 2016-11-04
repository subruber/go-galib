/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib selectors
*/

package ga

import (
	//	"log"
	"math/rand"
	"sort"
)

type GASelector interface {
	// Select one from pop
	SelectOne(pop GAGenomes) GAGenome

	// String name of selector
	String() string
}

//This selector first selects selector.Contestants random GAGenomes
//from the population then selects one based on PElite chance.
//The best contestant has PElite chance of getting selected.
//The next best contestant has PElite^2 chance of getting selected and so on
type GATournamentSelector struct {
	PElite      float64
	Contestants int
}

func NewGATournamentSelector(pelite float64, contestants int) *GATournamentSelector {
	if pelite == 0 {
		return nil
	}
	return &GATournamentSelector{pelite, contestants}
}

func (s *GATournamentSelector) SelectOne(pop GAGenomes) GAGenome {
	if s.Contestants < 2 || s.PElite == 0 {
		panic("Contestants and PElite are not set")
	}
	g := make(GAGenomes, s.Contestants)
	l := len(pop)
	for i := 0; i < s.Contestants; i++ {
		g[i] = pop[rand.Intn(l)]
	}
	sort.Sort(g)
	r := rand.Float64()
	p := s.PElite
	for i := 0; i < s.Contestants-1; i++ {
		if r < p {
			//	log.Printf("Selected %v with probability %v < %v", i, r, p)
			return g[i]
		} else {
			//	log.Printf("NOT Selected %v with probability %v >= %v", i, r, p)
		}

		p = p + p*(float64(1)-s.PElite)
	}
	//	log.Printf("Selected %v (fallback)", s.Contestants-1)
	return g[s.Contestants-1]
}
func (s *GATournamentSelector) String() string {
	return "GATournamentSelector"
}
