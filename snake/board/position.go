package board

//Position 左上を(0,0)とした座標
type Position struct {
	x int //X座標
	y int //Y座標
}

func (p *Position) equal(ps *Position) bool {
	return p.x == ps.x && p.y == ps.y
}