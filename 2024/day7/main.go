package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/MountVesuvius/advent-of-code/utils"
)


var s = regexp.MustCompile(": | ")

func check(arr []int, target int) int {
    fmt.Println(target)
    for i := range len(arr)-1 {
        s := make([]int, len(arr)-1)
        s[i] = 1
    fmt.Println(s)
    } 
    return 0 
}


func convert(arr []string) []int {
    s := make([]int, len(arr))
    for i := range arr { s[i] = utils.StrToNum(arr[i]) }
    return s
}

func main() {
    s := ""
    spaces := 3
    opts := []string{"+", "*"}
    a := int(math.Pow(float64(len(opts)), float64(spaces)))
    for range a {
        s += opts[0]
    }
    fmt.Println(s)
    // utils.GetInput(2024, 7)
    // data := utils.FileByLines("./sample")

    // for i := range data {
    //     a := convert(s.Split(data[i], -1))
    //     check(a[1:], a[0])
    // }
}
