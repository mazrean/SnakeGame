package board

import "math"

//Position 左上を(0,0)とした座標
type Position struct {
	X int //X座標
	Y int //Y座標
}

// ManhattanDistance マンハッタン距離を求める関数
func (p *Position) ManhattanDistance(ps *Position, d Direction) int {
	dx := p.X - ps.X
	dy := p.Y - ps.Y
	switch d {
	case UP:
		dy++
	case DOWN:
		dy--
	case RIGHT:
		dx--
	case LEFT:
		dx++
	}
	return int(math.Abs(float64(dx)) + math.Abs(float64(dy)))
}

func (p *Position) equal(ps *Position) bool {
	return p.X == ps.X && p.Y == ps.Y
}