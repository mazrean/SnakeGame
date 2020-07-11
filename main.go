package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/mazrean/SnakeGame/snake"
	"github.com/mazrean/SnakeGame/snake/board"
	"github.com/mazrean/SnakeGame/snake/measure"
)

func main() {

	flag.Parse()
	mode := flag.Arg(0)
	switch mode {
	case "solve":
		state, err := snake.ReadSnake()
		if err != nil {
			panic(fmt.Errorf("Read Initial State Error: %w", err))
		}

		searchType := flag.Arg(1)

		cpuf, err := os.Create("cpu.prof")
		if err != nil {
			panic(fmt.Errorf("CPU File Create Error: %w", err))
		}
		defer cpuf.Close()

		memf, err := os.Create("mem.prof")
		if err != nil {
			panic(fmt.Errorf("Memory File Create Error: %w", err))
		}

		defer func() {
			pprof.Lookup("heap").WriteTo(memf, 0)
			memf.Close()
		}()

		pprof.StartCPUProfile(cpuf)
		defer pprof.StopCPUProfile()

		directions, err := snake.SingleSnake(searchType, state)
		if err != nil {
			panic(fmt.Errorf("Search Answer Error: %w", err))
		}

		fmt.Println(directions)
	case "measure":
		snakeLen, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			panic(fmt.Errorf("Snake Length Convert Error: %w", err))
		}

		n, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			panic(fmt.Errorf("Board Convert Error: %w", err))
		}

		types := map[string]int{
			"A*": 10000,
			"iddfs": 10,
			"bfs": 10,
		}
		state, err := measure.Generate(snakeLen, n)
		if err != nil {
			panic(fmt.Errorf("Generate Error: %w", err))
		}

		var l int
		for k,v := range types {
			l, err = Measure(state, k, v)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(state, l)
	default:
		panic("Invalid Arguments")
	}
	
}

func Measure(state *board.State, v string, t int) (int, error) {
	cpuf, err := os.Create("cpu_" + v + ".prof")
	if err != nil {
		return 0, fmt.Errorf("CPU File Create Error: %w", err)
	}
	defer cpuf.Close()

	memf, err := os.Create("mem_" + v + ".prof")
	if err != nil {
		return 0, fmt.Errorf("Memory File Create Error: %w", err)
	}

	defer func() {
		pprof.Lookup("heap").WriteTo(memf, 0)
		memf.Close()
	}()

	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	var directions *[]board.Direction
	for i := 0; i < t; i++ {
		directions, err = snake.SingleSnake(v, state)
		if err != nil {
			panic(fmt.Errorf("Search Answer Error: %w", err))
		}
	}

	l := len(*directions)

	return l, nil
}