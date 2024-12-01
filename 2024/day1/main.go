package main

import (
	"fmt"
	"sort"

	"github.com/MountVesuvius/advent-of-code/utils"
)

func part2() {
    data := utils.FileSplit("./input", "   ")
    left, right := []int{}, []int{}

    for _, v := range data {
        left = append(left, utils.StrToNum(v[0]))
        right = append(right, utils.StrToNum(v[1]))
    }

    // fill table
    table := make(map[int]int)
    for _, v := range left { table[v] = 0 }

    for _, v := range right {
        _, e := table[v]
        if e {
            table[v] += 1
        }
    }

    total := 0
    for _, val := range left {
        total += val * table[val]
    }

    fmt.Println(total)
}

func part1() {
    data := utils.FileSplit("./input", "   ")
    left, right := []int{}, []int{}

    // Split lists
    for _, v := range data {
        left = append(left, utils.StrToNum(v[0]))
        right = append(right, utils.StrToNum(v[1]))
    }

    sort.Ints(left)
    sort.Ints(right)

    total := 0
    for i := 0; i < len(left); i++ {
        if left[i] > right[i] {
            total += left[i] - right[i]
        } else {
            total += right[i] - left[i]
        }
    }
    fmt.Println(total)
}

func main() {
    part1()
    part2()
}
