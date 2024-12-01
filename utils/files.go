package utils

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
)

// FileByLines opens the given path to a file and creates an array of strings
// with each element being a single line of input
func FileByLines(path string) ([]string) {
    data, err := os.Open(path)
    check(err)
    defer data.Close()

    scanner := bufio.NewScanner(data)
    scanner.Split(bufio.ScanLines)
    
    lines := []string{}
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return lines
}

// FileByNums expects the file input to be a single number per line.
// It creates an array of ints from each line
func FileByNums(path string) ([]int) {
    data, err := os.Open(path)
    check(err)
    defer data.Close()

    scanner := bufio.NewScanner(data)
    scanner.Split(bufio.ScanLines)
    
    nums := []int{}

    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        check(err)
        nums = append(nums, val)
    }

    return nums 
}

// FileByChars reads a single char out at a time, creating an array of chars
func FileByChars(path string) ([]string){
    data, err := os.Open(path)
    check(err)

    reader := bufio.NewReader(data)

    chars := []string{}
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            break
        }
        chars = append(chars, string(char))
    }

    return chars
}

// FileSplit is the generic for splitting an input file using a specific regex on each line
func FileSplit(path string, pattern string) ([][]string) {
    data, err := os.Open(path)
    check(err)
    defer data.Close()

    scanner := bufio.NewScanner(data)
    scanner.Split(bufio.ScanLines)
    r, err := regexp.Compile(pattern)
    check(err)

    var lines [][]string
    for scanner.Scan() {
        lines = append(lines, r.Split(scanner.Text(), -1))
    }

    return lines
}

// FileCSV Shortcut for FileSplit with , separation
func FileCSV(path string) ([][]string) {
    return FileSplit(path, ",")
}

// FileTSV Shortcut for FileSplit with tab separation
func FileTSV(path string) ([][]string) {
    return FileSplit(path, "\t")
}
