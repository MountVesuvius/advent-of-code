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

func Pop[T any](slice []T, index int) ([]T) {
	if index < 0 || index >= len(slice) {
		return nil
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice 
}
