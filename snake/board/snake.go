package board

// Snake 蛇の型
type Snake []*Position

// Head 頭の値
func (s *Snake) Head() (*Position, error) {
	return (*s)[len(*s)-1], nil
}

// Search 蛇の体が位置pにあるか探索
func (s *Snake) Search(p *Position) (bool, error) {
	for _,v := range (*s)[1:] {
		if p.equal(v) {
			return true, nil
		}
	}

	return false, nil
}

// Move 蛇の移動
func (s *Snake) Move(p *Position) (*Snake, error) {
	snake := append((*s)[1:], p)

	return &snake, nil
}

func (s *Snake) equal(sn *Snake) (bool,error) {
	if len(*s) != len(*sn) {
		return false, nil
	}
	for i,v := range *s {
		if !v.equal((*sn)[i]) {
			return false, nil
		}
	}

	return true, nil
}