package main

import (
	"flag"
	"fmt"

	"github.com/mazrean/SnakeGame/snake"
)

func main() {
	flag.Parse()
	searchType := flag.Arg(0)

	state, err := snake.ReadSnake()
	if err != nil {
		panic(fmt.Errorf("Read Initial State Error: %w", err))
	}

	directions, err := snake.SingleSnake(searchType, state)
	if err != nil {
		panic(fmt.Errorf("Search Answer Error: %w", err))
	}

	fmt.Println(directions)
}