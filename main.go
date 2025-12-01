package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Krivoguzov-Vlad/aoc/aoc"
	aoc2024 "github.com/Krivoguzov-Vlad/aoc/aoc/2024"
	aoc2025 "github.com/Krivoguzov-Vlad/aoc/aoc/2025"
)

func main() {
	years := map[int]aoc.AOC{
		2024: aoc2024.AOC,
		2025: aoc2025.AOC,
	}

	args := NewArgs()
	if len(os.Args) > 1 {
		args.Day, _ = strconv.Atoi(os.Args[1])
	}
	if len(os.Args) > 2 {
		args.Year, _ = strconv.Atoi(os.Args[2])
	}

	solver := years[args.Year][args.Day]
	solver.ReadInput(input(args))
	fmt.Println(solver.Part1())
	fmt.Println(solver.Part2())
}

func input(args Args) io.Reader {
	if file, ok := os.LookupEnv("AOC_INPUT"); ok {
		return InputFromFile(file)
	}
	if session, ok := os.LookupEnv("AOC_SESSION"); ok {
		return FetchInput(args, session)
	}
	return os.Stdin
}

type Args struct {
	Day  int
	Year int
}

func NewArgs() Args {
	now := time.Now()
	return Args{
		Year: now.Year(),
		Day:  now.Day(),
	}
}

func FetchInput(args Args, session string) io.Reader {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", args.Year, args.Day)
	r, _ := http.NewRequest(http.MethodGet, url, nil)
	r.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(fmt.Errorf("getting input: %w", err))
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("getting input: %w", err))
	}
	if resp.StatusCode >= 300 {
		panic(fmt.Errorf("getting input: %s", data))
	}
	return bytes.NewReader(data)
}

func InputFromFile(path string) io.Reader {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("getting input: %w", err))
	}
	return bytes.NewReader(data)
}
