package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


var client = &http.Client{}

func GetInput(year int, day int) {
    _, err := os.Open("./input")
    if err == nil {
        fmt.Println("Input file already found")
        return
    }
    // Get AoC session
    err = godotenv.Load("../../.env")
    check(err)
    session := os.Getenv("SESSION")

    // Build request for input file
    req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
    check(err)
    req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

    // Send request
    resp, err := client.Do(req)
    check(err)
    defer resp.Body.Close()

    // Create input file
    body, err := io.ReadAll(resp.Body)
    check(err)
    os.WriteFile("./input", body, 0644)
}

