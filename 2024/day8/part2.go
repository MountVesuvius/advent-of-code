package main

import (
	"fmt"

	"github.com/MountVesuvius/advent-of-code/utils"
)

// list[y][x]
type Coord struct {
    y int
    x int
}

// whatever the distance is between the two points, that's how far the antinodes have to be
func line2(a1, a2 Coord, dist int) []Coord {
    dx := a2.x - a1.x
    dy := a2.y - a1.y

    points := []Coord{a1}
    for i := 1; i <= dist; i++ {
        newPoint := Coord{
            x: a1.x + dx*i,
            y: a1.y + dy*i,
        }
        points = append(points, newPoint)
    }

    for i := 1; i <= dist; i++ {
        newPoint := Coord{
            x: a1.x - dx*i,
            y: a1.y - dy*i,
        }
        points = append(points, newPoint)
    }

    return points
}


func main() {
    utils.GetInput(2024, 8)
    data := utils.FileSplit("./input", "")

    // group antennas 
    antennas := make(map[string][]Coord)
    for i := range data {
        for j := range data[i] {
            if data[i][j] != "." {
                antennas[data[i][j]] = append(antennas[data[i][j]], Coord{
                    y: i,
                    x: j,
                })
            }
        }
    }

    // loop every set of antennas
    antinodes := []Coord{}
    for _, points := range antennas {
        if len(points) < 2 { continue } // don't bother process antennas with no friends 
        for i := range points {
            // for j := range points {
            for j := i + 1; j < len(points); j++ {
                if points[i] != points[j] {
                    a := line2(points[i], points[j], len(data[0]))
                    antinodes = append(antinodes, a...)
                }
            }
        }
   }

    c := make(map[Coord]int)
    for _, point := range antinodes {
        if point.y < len(data) && point.y >= 0 && point.x < len(data[point.y]) && point.x >= 0 {
            c[point] += 1
            if data[point.y][point.x] == "." {
                data[point.y][point.x] = "#"
            }
        }
    }

    fmt.Println(len(c))
    // fmt.Println(c)
    // for i := range data {
    //     fmt.Println(data[i])
    // }


}
