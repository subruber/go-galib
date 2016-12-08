/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

This genome is a list of int of given alphabet,
which is for problems where genes are limited by genes list.

For example, for REAL genetic problem in biology.
For DNA sequence, the "genes" (nucleotide acid bases) are "A", "C", "G", "T", "N" ...,
which can be encoded as 1, 2, 3, 4, 5 ...
*/

package ga

import (
	"fmt"
	"math/rand"
)

// GAIntOfAlphabetGenome is a list of int of given alphabet,
// which is for problems where genes are limited by genes list.
//
// For example, for REAL genetic problem in biology.
// For DNA sequence, the "genes" (nucleotide acid bases) are "A", "C", "G", "T", "N" ...,
// which can be encoded as 1, 2, 3, 4, 5 ...
type GAIntOfAlphabetGenome struct {
	Gene                 []int
	GeneAlphabet         []int
	weightedGeneAlphabet []int
	GeneAlphabetWeight   []int
	geneAlphabetMap      map[int]struct{}
	score                float64
	hasscore             bool
	sfunc                func(ga *GAIntOfAlphabetGenome) float64
}

// NewIntOfAlphabetGenome create Genome which consits of a list of int of given alphabet,
// and weight of letters in alphabet is optional, which will affect the result
// of function Randomize()
func NewIntOfAlphabetGenome(alphabet []int, weight []int, i []int, sfunc func(ga *GAIntOfAlphabetGenome) float64) (*GAIntOfAlphabetGenome, error) {
	g := new(GAIntOfAlphabetGenome)

	if len(alphabet) < 2 {
		return nil, fmt.Errorf("size of alphabet should be >= 2")
	}

	if len(weight) > 0 {
		if len(weight) != len(alphabet) {
			return nil, fmt.Errorf("number of alphabet (%d) and weight (%d) do not match", len(alphabet), len(weight))
		}
	} else {
		weight = make([]int, len(alphabet))
		for i := range alphabet {
			weight[i] = 1
		}
	}

	g.GeneAlphabet = alphabet
	g.GeneAlphabetWeight = weight

	g.weightedGeneAlphabet = []int{}
	for i, w := range g.GeneAlphabetWeight {
		for j := 0; j < w; j++ {
			g.weightedGeneAlphabet = append(g.weightedGeneAlphabet, g.GeneAlphabet[i])
		}
	}

	g.geneAlphabetMap = make(map[int]struct{}, len(alphabet))
	for _, a := range alphabet {
		g.geneAlphabetMap[a] = struct{}{}
	}

	for _, v := range i {
		if _, ok := g.geneAlphabetMap[v]; !ok {
			return nil, fmt.Errorf("invalid gene(%d) for alphabet(%v)", v, g.geneAlphabetMap)
		}
	}

	g.Gene = i
	g.sfunc = sfunc
	return g, nil
}

// Partially mapped crossover.
func (a *GAIntOfAlphabetGenome) Crossover(bi GAGenome, p1, p2 int) (GAGenome, GAGenome) {
	ca := a.Copy().(*GAIntOfAlphabetGenome)
	b := bi.(*GAIntOfAlphabetGenome)
	cb := b.Copy().(*GAIntOfAlphabetGenome)
	copy(ca.Gene[p1:p2+1], b.Gene[p1:p2+1])
	copy(cb.Gene[p1:p2+1], a.Gene[p1:p2+1])
	ca.Reset()
	cb.Reset()
	return ca, cb
}

func (a *GAIntOfAlphabetGenome) Splice(bi GAGenome, from, to, length int) {
	b := bi.(*GAIntOfAlphabetGenome)
	copy(a.Gene[to:length+to], b.Gene[from:length+from])
	a.Reset()
}

func (g *GAIntOfAlphabetGenome) Valid() bool {
	return true
}

func (g *GAIntOfAlphabetGenome) Switch(x, y int) {
	g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]
	g.Reset()
}

func (g *GAIntOfAlphabetGenome) Randomize() {
	l := len(g.weightedGeneAlphabet)
	for i := 0; i < len(g.Gene); i++ {
		g.Gene[i] = g.weightedGeneAlphabet[rand.Intn(l)]
	}
	g.Reset()
}

func (g *GAIntOfAlphabetGenome) Copy() GAGenome {
	n := new(GAIntOfAlphabetGenome)
	n.GeneAlphabet = make([]int, len(g.GeneAlphabet))
	copy(n.GeneAlphabet, g.GeneAlphabet)

	n.GeneAlphabetWeight = make([]int, len(g.GeneAlphabetWeight))
	copy(n.GeneAlphabetWeight, g.GeneAlphabetWeight)

	n.weightedGeneAlphabet = make([]int, len(g.weightedGeneAlphabet))
	copy(n.weightedGeneAlphabet, g.weightedGeneAlphabet)

	n.Gene = make([]int, len(g.Gene))
	copy(n.Gene, g.Gene)

	n.sfunc = g.sfunc
	n.score = g.score
	n.hasscore = g.hasscore
	return n
}

func (g *GAIntOfAlphabetGenome) Len() int { return len(g.Gene) }

func (g *GAIntOfAlphabetGenome) Score() float64 {
	if !g.hasscore {
		g.score = g.sfunc(g)
		g.hasscore = true
	}
	return g.score
}

func (g *GAIntOfAlphabetGenome) Reset() { g.hasscore = false }

func (g *GAIntOfAlphabetGenome) String() string { return fmt.Sprintf("%v", g.Gene) }
