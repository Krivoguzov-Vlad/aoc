package aoc2025

import (
	"io"
	"maps"
	"strconv"
	"strings"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day11 struct {
	graph map[string][]string
}

func (d *Day11) ReadInput(r io.Reader) {
	d.graph = make(map[string][]string)
	for _, line := range input.MustReadList[string](r, "\n") {
		nodes := strings.Split(line, " ")
		node := strings.TrimRight(nodes[0], ":")
		for _, adj := range nodes[1:] {
			d.graph[node] = append(d.graph[node], adj)
		}
	}
}

func (d *Day11) Part1() string {
	return strconv.Itoa(d.pathCount(d.graph, "you", "out"))
}

func (d *Day11) Part2() string {
	pathCount := func(graph map[string][]string) int {
		return d.pathCount(graph, "svr", "out")
	}
	all := pathCount(d.graph)
	withoutFft := pathCount(graphWithout(d.graph, "fft"))
	withoutDac := pathCount(graphWithout(d.graph, "dac"))
	withoutDacFft := pathCount(graphWithout(d.graph, "dac", "fft"))
	fft := all - withoutFft
	dac := all - withoutDac
	dacfft := all - withoutDacFft
	return strconv.Itoa(fft + dac - dacfft)
}

func (d *Day11) pathCount(graph map[string][]string, start, target string) int {
	pathCounts := make(map[string]int)
	pathCounts[target] = 1
	d.topologicalSort(start, graph, make(map[string]bool), pathCounts)
	return pathCounts[start]
}

func (d *Day11) topologicalSort(
	currentNode string,
	graph map[string][]string,
	visited map[string]bool, // just in case, there must be no cycles in the graph
	pathCounts map[string]int,
) {
	visited[currentNode] = true
	defer func() {
		visited[currentNode] = false
	}()

	pathCount := 0
	for _, adj := range graph[currentNode] {
		if visited[adj] {
			continue
		}
		if _, ok := pathCounts[adj]; !ok {
			d.topologicalSort(adj, graph, visited, pathCounts)
		}
		pathCount += pathCounts[adj]
	}
	pathCounts[currentNode] = pathCount
}

func graphWithout(graph map[string][]string, nodes ...string) map[string][]string {
	newGraph := maps.Clone(graph)
	for _, node := range nodes {
		delete(newGraph, node)
	}
	return newGraph
}
