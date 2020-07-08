package snake

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	bd "github.com/mazrean/SnakeGame/snake/board"
	"github.com/mazrean/SnakeGame/snake/collection"
)

var sc = bufio.NewScanner(os.Stdin)

// ReadSnake 入力処理
func ReadSnake() (*bd.State, error) {
	sc.Split(bufio.ScanWords)

	width, err := nextInt()
	if err != nil {
		return nil, fmt.Errorf("Width Read Error: %w", err)
	}
	height, err := nextInt()
	if err != nil {
		return nil, fmt.Errorf("Height Read Error: %w", err)
	}
	board := &bd.Board{
		Width: width,
		Height: height,
	}

	snakeLen, err := nextInt()
	if err != nil {
		return nil, fmt.Errorf("Snake Length Read Error: %w", err)
	}

	goalX, err := nextInt()
	if err != nil {
		return nil, fmt.Errorf("Goal X Read Error: %w", err)
	}
	goalY, err := nextInt()
	if err != nil {
		return nil, fmt.Errorf("Goal Y Read Error: %w", err)
	}
	goal := &bd.Position{
		X: goalX,
		Y: goalY,
	}

	snakeQueue := collection.Queue{}
	// しっぽから頭へ順に蛇の位置
	for i := 0; i < snakeLen; i++ {
		bodyX, err := nextInt()
		if err != nil {
			return nil, fmt.Errorf("Goal X Read Error: %w", err)
		}

		bodyY, err := nextInt()
		if err != nil {
			return nil, fmt.Errorf("Goal Y Read Error: %w", err)
		}

		position := &bd.Position{
			X: bodyX,
			Y: bodyY,
		}

		err = snakeQueue.Push(position)
		if err != nil {
			return nil, fmt.Errorf("Snake Push Error: %w", err)
		}
	}

	state := &bd.State{
		Turn: 0,
		Directions: []bd.Direction{},
		Board: board,
		Goal: goal,
		Snake: bd.Snake(snakeQueue),
	}

	return state, nil
}

func nextInt() (int,error) {
	isEnd := sc.Scan()
	if !isEnd {
		return 0, errors.New("No Readable Value")
	}

    i, err := strconv.Atoi(sc.Text())
    if err != nil {
        return 0, fmt.Errorf("Combert To Number Error: %w", err)
	}

    return i, nil
}