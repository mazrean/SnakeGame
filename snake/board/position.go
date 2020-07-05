package board

//左上を(0,0)とした座標
type position struct {
	x uint //X座標
	y uint //Y座標
}

func (p *position) equal(ps *position) bool {
	return p.x == ps.x && p.y == ps.y
}