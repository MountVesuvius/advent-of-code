package utils

import "strconv"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func StrToNum(in string) (out int) {
    a, err := strconv.Atoi(in)
    if err != nil {
        panic(err)
    }
    return a
}

func Pop[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}
