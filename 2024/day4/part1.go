package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/MountVesuvius/advent-of-code/utils"
)

var r1 = regexp.MustCompile("XMAS")
var r2 = regexp.MustCompile("SAMX")

func matches(s string) (int) {
    return len(r1.FindAllString(s, -1)) + len(r2.FindAllString(s, -1))
}

func DiagTR(matrix []string) ([]string) {
	rows := len(matrix)
	cols := len(matrix[0])
    s := []string{}

	for col := cols - 1; col >= 0; col-- {
		s = append(s, Diag(matrix, 0, col, 1, -1))
	}

	for row := 1; row < rows; row++ {
		s = append(s, Diag(matrix, row, cols-1, 1, -1))
	}

    return s
}

func DiagTL(matrix []string) ([]string) {
	rows := len(matrix)
	cols := len(matrix[0])
    s := []string{}

	for col := 0; col < cols; col++ {
		s = append(s, Diag(matrix, 0, col, 1, 1))
	}

	for row := 1; row < rows; row++ {
		s = append(s, Diag(matrix, row, 0, 1, 1))
	}

    return s
}

func Diag(matrix []string, startRow, startCol, rowStep, colStep int) (string) {
	rows := len(matrix)
	cols := len(matrix[0])
	row, col := startRow, startCol
	var diagonal []string

	for row >= 0 && row < rows && col >= 0 && col < cols {
		diagonal = append(diagonal, string(matrix[row][col]))
		row += rowStep
		col += colStep
	}

	return strings.Join(diagonal, "")
}


func main() {
    data := utils.FileByLines("./input")
    total := 0

    // rows
    for _, v := range data {
        total += matches(v)
    }

    // cols
    for i := range data[0] {
        tmp := []string{}
        for j := range data {
            tmp = append(tmp, string(data[j][i]))
        }
        total += matches(strings.Join(tmp, ""))
    }

    // diags
    a := DiagTR(data)
    for _, v := range a {
        total += matches(v)
    }
    b := DiagTL(data)
    for _, v := range b {
        total += matches(v)
    }
    fmt.Println(total)

}
