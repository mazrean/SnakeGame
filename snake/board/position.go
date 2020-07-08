package board

//Position 左上を(0,0)とした座標
type Position struct {
	X int //X座標
	Y int //Y座標
}

func (p *Position) equal(ps *Position) bool {
	return p.X == ps.X && p.Y == ps.Y
}