package swarm

import (
	"fmt"
	"math/rand"
	"sync"
)

type Swarm struct{}

type Ans struct {
	ans float64
	res []int
}

func (s *Swarm) Solve(p Problem) []int {
	var wg sync.WaitGroup
	c := make(chan Ans, 50)

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go s.exec(p, c, &wg)
	}
	wg.Wait()
	close(c)

	best := 999999.0
	var cBest []int
	for val := range c {
		if best > val.ans {
			best = val.ans
			cBest = val.res
		}
	}
	return cBest
}

func (s *Swarm) exec(p Problem, c chan Ans, wg *sync.WaitGroup) {
	fmt.Println("Got here")
	sol := make([]int, 0)
	for i := range p.Nodes {
		sol = append(sol, i)
	}
	rand.Shuffle(len(sol), func(i, j int) { sol[i], sol[j] = sol[j], sol[i] })
	c <- Ans{p.Score(sol), sol}
	wg.Done()
	fmt.Println("DONE 1")
}
