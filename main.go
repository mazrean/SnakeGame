package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/mazrean/SnakeGame/snake"
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
	case "mesure":
		snakeLen, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			panic(fmt.Errorf("Snake Length Convert Error: %w", err))
		}

		n, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			panic(fmt.Errorf("Board Convert Error: %w", err))
		}

		types := []string{"A*", "iddfs", "bfs"}
		state, err := measure.Generate(snakeLen, n)
		if err != nil {
			panic(fmt.Errorf("Generate Error: %w", err))
		}

		for _,v := range types {
			err := func() error {
				cpuf, err := os.Create("cpu_" + v + ".prof")
				if err != nil {
					return fmt.Errorf("CPU File Create Error: %w", err)
				}
				defer cpuf.Close()

				memf, err := os.Create("mem_" + v + ".prof")
				if err != nil {
					return fmt.Errorf("Memory File Create Error: %w", err)
				}

				defer func() {
					pprof.Lookup("heap").WriteTo(memf, 0)
					memf.Close()
				}()

				pprof.StartCPUProfile(cpuf)
				defer pprof.StopCPUProfile()

				directions, err := snake.SingleSnake(v, state)
				if err != nil {
					panic(fmt.Errorf("Search Answer Error: %w", err))
				}

				fmt.Println(directions)

				return nil
			}()
			if err != nil {
				panic(err)
			}
		}
	default:
		panic("Invalid Arguments")
	}
	
}