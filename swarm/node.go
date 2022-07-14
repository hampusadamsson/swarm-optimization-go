package swarm

import "math"

type Node struct {
	X float64
	Y float64
}

func (n *Node) dist(n2 *Node) float64 {
	return math.Abs(n.X-n2.X) + math.Abs(n.Y-n2.Y)
}
