package collection

import (
	"errors"
	"fmt"
)

// Queue キューの構造体
type Queue []interface{}

// Push 値の追加
func (q *Queue) Push(v interface{}) error {
	*q = append(*q, v)

	return nil
}

// Pop 値の取り出し&削除
func (q *Queue) Pop() (interface{}, error) {
	v, err := q.Peek()
	if err != nil {
		return nil, err
	}

	*q = (*q)[1:]

	return v, nil
}

// Peek 値の取り出し
func (q *Queue) Peek() (interface{}, error) {
	empty, err := q.Empty()
	if err != nil {
		return nil, fmt.Errorf("Check Empty Error: %w", err)
	}
	if empty {
		return nil, errors.New("Empty Queue")
	}

	return (*q)[0], nil
}

// Size サイズ
func (q *Queue) Size() (int, error) {
	return len(*q), nil
}

// Empty 空か
func (q *Queue) Empty() (bool, error) {
	return len(*q) ==0, nil
}