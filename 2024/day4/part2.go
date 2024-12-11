package main

import (
	"fmt"
	"regexp"
	"strings"

	// "strings"

	"github.com/MountVesuvius/advent-of-code/utils"
)

var ends = regexp.MustCompile("M.S")
// var ends2 = regexp.MustCompile("S.M")
var mids = regexp.MustCompile(".A.")

func chunk(data []string) int {
    total := 0
    for i := 2; i < len(data); i++ {
        block := data[i-2:i+1]
        total += check(block)
    }
    return total
}

func check(block []string) (int) {
    total := 0
        for j := 2; j < len(block[0]); j++ {
            top := block[0][j-2:j+1]
            mid := block[1][j-2:j+1]
            bottom := block[2][j-2:j+1]
            a := len(ends.FindString(top))
            b := len(ends.FindString(bottom))
            c := len(mids.FindString(mid))
            if (a + b + c) == 9 {
                total++
            }
        }
        return total
}

func rotate(matrix []string) []string {
    rows := len(matrix)
    if rows == 0 {
        return nil
    }
    cols := len(matrix[0])
    
    s := make([]string, cols)

    for i := 0; i < cols; i++ {
        tmp := []string{}
        for j := rows - 1; j >= 0; j-- {
            tmp = append(tmp, string(matrix[j][i]))
        }
        s[i] = strings.Join(tmp, "")
    }
    return s
}

func main() {
    data := utils.FileByLines("./input")

    r1 := rotate(data)
    r2 := rotate(r1)
    r3 := rotate(r2)

    a := chunk(data) + chunk(r1) + chunk(r2) + chunk(r3)
    fmt.Println(a)

}
