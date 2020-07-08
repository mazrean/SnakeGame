package board

import (
	"errors"

	"github.com/mazrean/SnakeGame/snake/collection"
)

// Snake 蛇の型
type Snake collection.Queue

// Head 頭の値
func (s *Snake) Head() (*Position, error) {
	interfaceHead := (*s)[len(*s)-1]

	head, ok := interfaceHead.(*Position)
	if !ok {
		return nil, errors.New("Unexpected Parse Type Error")
	}

	return head, nil
}

// Search 蛇の体が位置pにあるか探索
func (s *Snake) Search(p *Position) (bool, error) {
	for _,interfaceVal := range (*s)[1:] {
		val, ok := interfaceVal.(*Position)
		if !ok {
			return false, errors.New("Unexpected Parse Type Error")
		}

		if p.equal(val) {
			return true, nil
		}
	}

	return false, nil
}

func (s *Snake) equal(sn *Snake) (bool,error) {
	if len(*s) != len(*sn) {
		return false, nil
	}
	for i,v := range *s {
		p, ok := v.(*Position)
		if !ok {
			return false, errors.New("Unexpected Type Parse Error")
		}

		interfacePs := (*sn)[i]
		ps, ok := interfacePs.(*Position)
		if !ok {
			return false, errors.New("Unexpected Type Parse Error")
		}

		if !p.equal(ps) {
			return false, nil
		}
	}

	return true, nil
}