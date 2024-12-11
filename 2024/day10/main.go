package main

import (
	"fmt"
	"time"

	"github.com/MountVesuvius/advent-of-code/utils"
)

type Pos struct {
    row, col int
}

type Edge struct {
    From, To Pos
    Value int
}

type Graph struct {
    Nodes map[Pos]int
    Edges map[Pos][]Edge
}

func newGraph() *Graph {
    return &Graph{ make(map[Pos]int), make(map[Pos][]Edge)}
}

func (g *Graph) AddNode(pos Pos, id int) {
    g.Nodes[pos] = id
}

func (g *Graph) AddEdge(from, to Pos, value int) {
    g.Edges[from] = append(g.Edges[from], Edge{from, to, value})
}

var directions = [][]int{
    {-1, 0}, // up
    {1, 0},  // down
    {0, -1},  // left
    {0, 1},  // right
}

func checkBounds(val, top int) bool {
    return val < top && val >= 0
}

func buildGraph(matrix[][]int) (*Graph, []Pos) {
    graph := newGraph()

    rows := len(matrix)
    cols := len(matrix[0])

    trailheads := []Pos{}
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            node := Pos{row, col}
            if matrix[row][col] == 0 {
                trailheads = append(trailheads, node)
            }
            graph.AddNode(node, matrix[row][col])

            for _, dir := range directions {
                newRow := row + dir[0]
                newCol := col + dir[1]

                if checkBounds(newRow, rows) && checkBounds(newCol, cols) {
                    if matrix[newRow][newCol] == matrix[row][col] + 1 {
                        graph.AddEdge(node, Pos{newRow, newCol}, 1)
                    }
                }
            }
        }
    }
    return graph, trailheads 
}

func findConnectedNodes(graph *Graph, start Pos) []Pos {
    visited := make(map[Pos]bool)
    connectedNodes := []Pos{}
    stack := []Pos{start}

    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if !visited[current] {
            visited[current] = true
            connectedNodes = append(connectedNodes, current)

            for _, edge := range graph.Edges[current] {
                if !visited[edge.To] {
                    stack = append(stack, edge.To)
                }
            }
        }
    }

    return connectedNodes
}

func distinctPaths(graph *Graph, start, target Pos) int {
    if start == target {
        return 1
    }
    pathCount := 0
    for _, edge := range graph.Edges[start] {
        pathCount += distinctPaths(graph, edge.To, target)
    }
    return pathCount
}

func main() {
    utils.GetInput(2024, 10)
    data := utils.FileSplit("./input", "")

    start := time.Now()
    // parse for just ints
    topography := [][]int{}
    for i := range data {
        topoRow := []int{}
        for j := range data[i] {
            topoRow = append(topoRow, utils.StrToNum(data[i][j]))
        }
        topography = append(topography, topoRow)
    }

    graph, trailheads := buildGraph(topography)

    // this works because i don't allow sub graphs to be connected to nodes that aren't part of a path
    // as a result you can just count the nodes == 9 as they are guaranteed to be on a path
    total := 0
    distict := 0
    for i := range trailheads {
        subtotal := 0
        for _, node := range findConnectedNodes(graph, trailheads[i]) {
            if graph.Nodes[node] == 9 {
                distict += distinctPaths(graph, trailheads[i], node)
                subtotal++
            }
        }
        total += subtotal
    }
    fmt.Println(time.Since(start))

    fmt.Println(total)
    fmt.Println(distict)

}
