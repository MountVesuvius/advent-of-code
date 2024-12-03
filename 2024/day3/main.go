package main

import (
	"fmt"
	"regexp"

	"github.com/MountVesuvius/advent-of-code/utils"
)

func part1() {
    data := utils.FileByLines("./input")
    r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
    d, _ := regexp.Compile(`\d+`)

    total := 0
    for _, i := range data {
        for _, j := range r.FindAllString(i, -1) {
            nums := d.FindAllString(j, -1) 
            total += utils.StrToNum(nums[0]) * utils.StrToNum(nums[1])
        }
    }
    fmt.Println(total)
}

func part2() {
    data := utils.FileByLines("./input")
    r, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
    d, _ := regexp.Compile(`\d+`)

    total := 0
    active := true
    for _, i := range data {
        for _, i := range r.FindAllString(i, -1) {
            if i == "don't()" {
                active = false
                continue
            }
            if i == "do()" {
                active = true 
                continue
            }
            if active {
                nums := d.FindAllString(i, -1) 
                total += utils.StrToNum(nums[0]) * utils.StrToNum(nums[1])
            }
        }
    }
    fmt.Println(total)
}

func main() {
    part1()
    part2()
}
