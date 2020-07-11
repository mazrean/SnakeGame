package collection

import (
	"errors"
	"fmt"
)

// Stack スタックの構造体
type Stack []*Node

// Push 値の追加
func (s *Stack) Push(v *Node) error {
	*s = append(*s, v)

	return nil
}

// Pop 値の取り出し&削除
func (s *Stack) Pop() (*Node, error) {
	v, err := s.Peek()
	if err != nil {
		return nil, err
	}

	*s = (*s)[1:]

	return v, nil
}

// Peek 値の取り出し
func (s *Stack) Peek() (*Node, error) {
	empty, err := s.Empty()
	if err != nil {
		return nil, fmt.Errorf("Check Empty Error: %w", err)
	}
	if empty {
		return nil, errors.New("Empty Queue")
	}

	return (*s)[len(*s)-1], nil
}

// Size サイズ
func (s *Stack) Size() (int, error) {
	return len(*s), nil
}

// Empty 空か
func (s *Stack) Empty() (bool, error) {
	return len(*s) ==0, nil
}

func (s *Stack) String() string {
	str := ""
	for i,v := range *s {
		str += fmt.Sprintf("%d:%s\n%+v\n", i, v.Direction, v.State)
	}

	return str
}