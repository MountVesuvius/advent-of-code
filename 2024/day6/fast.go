package main

import (
	"fmt"
	"time"

	"github.com/MountVesuvius/advent-of-code/utils"
)

func run(guard []int, lookupTable map[int][]int) {
    fmt.Println("wehlp")
}

func main() {
    start := time.Now()
    data := utils.FileByLines("./input")

    // setup
    rowLookup := make(map[int][]int)
    colLookup := make(map[int][]int)
    guard := []int{0, 0, 0} // 0 up, 1 right, 2 bottom, 3 left
    for i := range data {
        for j := range data[i] {
            if data[i][j] == '#' {
                rowLookup[i] = append(rowLookup[i], j)
                colLookup[j] = append(colLookup[j], i)
            }
            if data[i][j] == '^' {
                guard[0] = i
                guard[1] = j
            }
        }
    }
    run(guard, rowLookup)
    fmt.Println(time.Since(start))

    // fmt.Println(rowLookup, colLookup)
    // fmt.Println(guard)
}
