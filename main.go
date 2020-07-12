package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"strconv"
	"syscall"
	"time"

	"github.com/mazrean/SnakeGame/snake"
	"github.com/mazrean/SnakeGame/snake/board"
	"github.com/mazrean/SnakeGame/snake/measure"
)

type resultStruct struct {
	depth int
	node int64
	snake int
	width int
	time int64
}

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

		start := time.Now();
		directions, node, err := snake.SingleSnake(searchType, state)
		if err != nil {
			panic(fmt.Errorf("Search Answer Error: %w", err))
		}
		fmt.Println(node)
		end := time.Now();
		fmt.Printf("%dms %dμs\n",end.Sub(start).Milliseconds(),end.Sub(start).Microseconds())

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
			"A*": 1,
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
	case "gan":
		as := []resultStruct{}
		iddfss := []resultStruct{}
		bfss := []resultStruct{}
		resultMap := map[string]*[]resultStruct{
			"A*": &as,
			"iddfs": &iddfss,
			"bfs": &bfss,
		}
		sigc := make(chan os.Signal)
		go func(){
			signal.Notify(sigc, syscall.SIGTERM)
		}()
		go func(){
			<-sigc
			output(resultMap)
		}()

		limMap := map[string]int{
			"A*": 2,
			"iddfs": 2,
			"bfs": 2,
		}
		ulim := 2
		for width := 2; width < 50; width++ {
			lim := width*2
			for snakeLen := lim; snakeLen >= ulim; snakeLen-- {
				state, err := measure.Generate(snakeLen, width)
				if err != nil {
					continue
				}
				for t, l := range limMap {
					if l >= snakeLen {
						break
					}
					start := time.Now()
					directions, node, err := snake.SingleSnake(t, state)
					end := time.Now()
					time := end.Sub(start).Microseconds()
					if err != nil {
						break
					}
					res := resultStruct{
						depth: len(*directions),
						node: node,
						snake: snakeLen,
						width: width,
						time: time,
					}
					*resultMap[t] = append(*resultMap[t], res)
					fmt.Printf("%#v\n", res)
					if res.time > 1000000 || res.depth >= 20 {
						ulim = snakeLen
						break
					}
				}
			}
		}
		output(resultMap)
	default:
		panic("Invalid Arguments")
	}
	
}

func output(resultMap map[string]*[]resultStruct) {
	for t, arr := range resultMap {
		file, err := os.OpenFile(t + ".csv", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			continue
		}
		defer file.Close()

		for _,v := range *arr {
			line := fmt.Sprintf("%d, %d, %d, %d, %d\n", v.depth, v.node, v.snake, v.width, v.time)
			file.WriteString(line)
		}
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
	var node int64
	for i := 0; i < t; i++ {
		directions, node, err = snake.SingleSnake(v, state)
		if err != nil {
			panic(fmt.Errorf("Search Answer Error: %w", err))
		}
		fmt.Println(node)
	}

	l := len(*directions)

	return l, nil
}