package main

import (
	"fmt"
	"math"

	"github.com/MountVesuvius/advent-of-code/utils"
)

func nuclearTest(list []string) (bool) {
    swing := make(map[bool]int)
    for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
        temp := utils.StrToNum(list[i]) - utils.StrToNum(list[j])
        swing[temp >= 0] += 1

        if len(swing) > 1 {
            return false
        }

        diff := math.Abs(float64(temp))
        if diff > 3 || diff == 0 {
            return false
        }
    }
    return true
}


func part2() {
    data := utils.FileSplit("./input", " ")

    c := 0
    for _, list := range data {
        if nuclearTest(list) {
            c++
            continue
        }

        test := make([]string, len(list))
        for i := range list {
            copy(test, list)
            test := utils.Pop(test, i)
            if nuclearTest(test) {
                c++
                break 
            }
        }

    }
    fmt.Println(c)
}

func part1() {
    data := utils.FileSplit("./input", " ")

    c := 0
    for _, list := range data { if nuclearTest(list) { c++ } }
    fmt.Println(c)
}


func main() {
    part1()
    part2()
}
