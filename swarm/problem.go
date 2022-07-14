package swarm

import (
	"math/rand"
)

type Problem struct {
	Nodes []Node
}

func (p *Problem) Score(path []int) float64 {
	dist := 0.0
	for i := range path {
		node := p.Nodes[path[i]]
		nextNode := p.Nodes[path[(i+1)%len(p.Nodes)]]
		dist += node.dist(&nextNode)
	}
	return dist
}

func (p *Problem) Normalize() {
	p.scaleX()
	p.scaleY()
}

func (p *Problem) scaleX() {
	max := 0.0
	for _, n := range p.Nodes {
		if n.X > max {
			max = n.X
		}
	}
	for i := range p.Nodes {
		p.Nodes[i].X = p.Nodes[i].X / max
	}
}

func (p *Problem) scaleY() {
	max := 0.0
	for _, n := range p.Nodes {
		if n.Y > max {
			max = n.Y
		}
	}
	for i := range p.Nodes {
		p.Nodes[i].Y = p.Nodes[i].Y / max
	}
}

// RandomProblem - retrieve an example problem
func RandomProblem(nC int) Problem {
	ns := make([]Node, 0)
	for i := 0; i < nC; i++ {
		ns = append(ns, Node{rand.Float64(), rand.Float64()})
	}
	return Problem{
		Nodes: ns,
	}
}
