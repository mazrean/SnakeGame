package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/mazrean/SnakeGame/snake"
)

func main() {
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