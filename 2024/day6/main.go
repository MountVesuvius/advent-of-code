package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/MountVesuvius/advent-of-code/utils"
)

var r = regexp.MustCompile("")
var g = regexp.MustCompile(`\^|v|<|>`)
var p = regexp.MustCompile("X")

func step(char string) (int, int) {
    switch char {
        case "^":
            return -1, 0
        case "v":
            return 1, 0
        case ">":
            return 0, 1
        case "<":
            return 0, -1
    }
    return 0, 0
}

func display(grid [][]string) {
    for i := range grid {
        fmt.Println(grid[i])
    }
    fmt.Println()
}


func turn(char string) string {
    switch char {
        case "^":
            return ">"
        case ">":
            return "v"
        case "v":
            return "<" 
        case "<":
            return "^"
        }
    return "H"
}

func move(grid [][]string, guard []int, limit int) int {
    m := 0
    escape := 1
    for m < limit {
        m++
        // finish state
        dir := grid[guard[0]][guard[1]]
        a, b := step(dir)
        if guard[0]+a >= len(grid) || guard[0]+a < 0 || guard[1]+b >= len(grid[0]) || guard[1]+b < 0 {
            escape = 0
            break
        }

        next := grid[guard[0] + a][guard[1] + b]

        // move on board
        if next != "#" && next != "O" {
            grid[guard[0]][guard[1]] = "X"
            guard[0] += a
            guard[1] += b
            grid[guard[0]][guard[1]] = dir
        } else {
            x := grid[guard[0]][guard[1]]
            grid[guard[0]][guard[1]] = turn(x)
        }
    }
    return escape
}

func part1() {
    data := utils.FileByLines("./sample")

    // setup
    grid := [][]string{}
    guard := []int{0, 0}
    for i, v := range data {
        line := r.Split(v, -1)
        grid = append(grid, line)
        for j := range line {
            if len(g.FindAllString(line[j], -1)) > 0 {
                guard[0] = i
                guard[1] = j
            }
        }
    }
    move(grid, guard, 16_900)

    total := 0
    for i := range grid {
        a := strings.Join(grid[i], "")
        total += len(p.FindAllString(a, -1))
    }
    fmt.Println(total+1)
}

func part2() {
    data := utils.FileByLines("./input")

    // setup
    grid := [][]string{}
    guard := []int{0, 0}
    for i, v := range data {
        line := r.Split(v, -1)
        grid = append(grid, line)
        for j := range line {
            if len(g.FindAllString(line[j], -1)) > 0 {
                guard[0] = i
                guard[1] = j
            }
        }
    }

    hold := []int{0, 0}
    copy(hold, guard)
    c := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == "." || grid[i][j] == "X" {
                // copy grid
                tempGrid := make([][]string, len(grid))
                for k := range grid {
                    tempGrid[k] = append([]string{}, grid[k]...)
                }

                tempGrid[i][j] = "O"
                copy(guard, hold)
                c += move(tempGrid, guard, 16_900)

                // reset
                tempGrid[i][j] = "."
            }
        }
    }
    fmt.Println("final count", c)
}

func m() {
    // utils.GetInput(2024, 6)
    start := time.Now()
    part1()
    fmt.Println(time.Since(start))
    part2()
    fmt.Println(time.Since(start))
}
