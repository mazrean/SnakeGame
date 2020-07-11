package measure

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mazrean/SnakeGame/snake/board"
)

// Generate 盤面の自動生成
func Generate(snakeLen int, n int) (*board.State, error) {
	rand.Seed(time.Now().UnixNano())

	nulDirections := []board.Direction{}

	snake := make(board.Snake, 0, snakeLen)
	snake = append(snake, &board.Position{
		X: rand.Intn(n),
		Y: rand.Intn(n),
	})

	for i := 1; i < snakeLen; i++ {
		positions := []*board.Position{
			{
				X: snake[i-1].X,
				Y: snake[i-1].Y-1,
			},
			{
				X: snake[i-1].X,
				Y: snake[i-1].Y+1,
			},
			{
				X: snake[i-1].X+1,
				Y: snake[i-1].Y,
			},
			{
				X: snake[i-1].X-1,
				Y: snake[i-1].Y,
			},
		}

		pos := make([]*board.Position, 0, 4)
		for _,v := range positions {
			isThere, err := snake.Search(v)
			if err != nil {
				return nil, fmt.Errorf("Snake Search Error: %w", err)
			}

			if v.X < 0 || v.X >= n || v.Y < 0 || v.Y >= n || isThere {
				continue
			}
			pos = append(pos, v)
		}

		if len(pos) == 0 {
			return Generate(snakeLen, n)
		}

		position := pos[rand.Intn(len(pos))]
		snake = append(snake, position)
	}

	state := &board.State{
		Turn: 0,
		Directions: &nulDirections,
		Board: &board.Board{
			Width: n,
			Height: n,
		},
		Goal: &board.Position{
			X: rand.Intn(n),
			Y: rand.Intn(n),
		},
		Snake: &snake,
	}

	return state, nil
}
