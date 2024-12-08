package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/MountVesuvius/advent-of-code/utils"
)

var s = regexp.MustCompile(`\d+`)
var c = regexp.MustCompile(`,`)

func contains(arr []int, target int) int {
    for _, v := range arr {
        if target == v {
            return v
        }
    }
    return -1
}

func part1() {
    utils.GetInput(2024, 5)
    data := utils.FileByLines("./input")

    orderingRules := make(map[int][]int)
    from := 1
    for _, v := range data {
        a := s.FindAllString(v, -1)
        if len(a) != 2 { break }
        before := utils.StrToNum(a[0])
        after := utils.StrToNum(a[1])
        orderingRules[before] = append(orderingRules[before], after)
        from++
    }

    // loop through page orders
    total := 0
    for _, v := range data[from:] {
        a := c.Split(v, -1)
        // each one pull the value
        toggle := true
        for i := range a {
            val := utils.StrToNum(a[i])
            for j := i+1; j < len(a); j++ {
                x := utils.StrToNum(a[j])
                b := contains(orderingRules[val], x)
                if b == -1 {
                    toggle = false
                }
            }
        }
        if (toggle) {
            total += utils.StrToNum(a[len(a)/2])
        }
    }
    fmt.Println(total)
}

func reorderList(orderingRules map[int][]int, list []string) []string {
	sort.SliceStable(list, func(i, j int) bool {
        a := utils.StrToNum(list[i])
        b := utils.StrToNum(list[j])
		return contains(orderingRules[a], b) != -1
	})
	return list
}


func main() {
    utils.GetInput(2024, 5)
    data := utils.FileByLines("./input")

    orderingRules := make(map[int][]int)
    from := 1
    for _, v := range data {
        a := s.FindAllString(v, -1)
        if len(a) != 2 { break }
        before := utils.StrToNum(a[0])
        after := utils.StrToNum(a[1])
        orderingRules[before] = append(orderingRules[before], after)
        from++
    }

    // loop through page orders
    arr := [][]string{}
    for _, v := range data[from:] {
        a := c.Split(v, -1)
        // each one pull the value
        toggle := false 
        for i := range a {
            val := utils.StrToNum(a[i])
            for j := i+1; j < len(a); j++ {
                x := utils.StrToNum(a[j])
                b := contains(orderingRules[val], x)
                if b == -1 {
                    toggle = true 
                }
            }
        }
        if (toggle) {
            arr = append(arr, a)
        }
    }

    fmt.Println(orderingRules)
    for _, list := range arr {
        for i, v := range list {
            val := utils.StrToNum(v)
            for j := i+1; j < len(list); j++ {
                x := utils.StrToNum(list[j])
                if contains(orderingRules[val], x) == -1 {
                    reorderList(orderingRules, list)
                }
            }

            fmt.Println(list)
        }
        fmt.Println()
    }

    total := 0
    for _, list := range arr {
        total += utils.StrToNum(list[len(list)/2])
    }

    fmt.Println(total)
}


