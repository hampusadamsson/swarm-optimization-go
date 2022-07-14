package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hampusadamsson/swarm-optimization-go/swarm"
)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	X = 960
	Y = 720
)

func init() {
	emptyImage.Fill(color.White)
}

type Game struct {
	problem  swarm.Problem
	solution []int
	best     float64
	sol      swarm.Swarm
}

func (g *Game) DrawLine(dst *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color) {
	x1 = x1 * float64(X-100)
	x2 = x2 * float64(X-100)
	y1 = y1 * float64(Y-100)
	y2 = y2 * float64(Y-100)

	length := math.Hypot(x2-x1, y2-y1)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(length, 1)
	op.GeoM.Rotate(math.Atan2(y2-y1, x2-x1))
	op.GeoM.Translate(x1, y1)
	op.ColorM.ScaleWithColor(clr)
	dst.DrawImage(emptySubImage, op)
}

func (g *Game) DrawCircle(dst *ebiten.Image, cx, cy, r float64, clr color.Color) {
	cx = cx * float64(X-100)
	cy = cy * float64(Y-100)

	var path vector.Path
	rd, gg, b, a := clr.RGBA()
	path.Arc(float32(cx), float32(cy), float32(r), 0, 2*math.Pi, vector.Clockwise)
	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vertices {
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(rd) / 0xffff
		vertices[i].ColorG = float32(gg) / 0xffff
		vertices[i].ColorB = float32(b) / 0xffff
		vertices[i].ColorA = float32(a) / 0xffff
	}
	dst.DrawTriangles(vertices, indices, emptySubImage, nil)
}

func (g *Game) Update() error {
	// Calc path
	sol := g.sol.Solve(g.problem)
	dist := g.problem.Score(sol)
	if dist < g.best {
		fmt.Println(dist, sol)
		g.best = dist
		g.solution = sol
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw cities
	for _, node := range g.problem.Nodes {
		g.DrawCircle(screen, node.X, node.Y, 10, color.White)
	}
	// Draw path
	for i := range g.solution {
		cur := g.solution[i]
		next := g.solution[(i+1)%len(g.solution)]
		node := g.problem.Nodes[cur]
		nextNode := g.problem.Nodes[next]
		g.DrawCircle(screen, node.X, node.Y, 10, color.White)
		g.DrawLine(screen, node.X, node.Y, nextNode.X, nextNode.Y, color.White)
	}
	// print status
	s := fmt.Sprintf("%f", g.best)
	ebitenutil.DebugPrint(screen, s)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return X, Y
}

func main() {
	ebiten.SetWindowSize(X, Y)
	ebiten.SetWindowTitle("Hello, World!")

	prob := swarm.RandomProblem(25)
	prob.Normalize()

	s := swarm.Swarm{}

	g := &Game{
		problem:  prob,
		sol:      s,
		solution: s.Solve(prob),
		best:     9999999999,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
