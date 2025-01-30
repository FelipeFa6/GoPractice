package main

import (
    "fmt"
)

// implementa una función que cuente la
// cantidad de veces que aparece cada palabra en
// un string

func main() {
    input1 := []string{
        "T1", "T1", "T1", "T1",     // T1 = 4 times
        "Fish",                     // Fish = 1 time
        "COW", "COW",               // COW = 2 times
        "cow", "cow", "cow", "cow", // COW = 3 times
    }

    response := wordCounter(input1)

    for word, count := range response {
        fmt.Printf("%s: %d veces\n", word, count)
    }
}

func wordCounter(arr []string) map[string]int {
    countMap := make(map[string]int)

    for _, word := range arr {
        countMap[word]++
    }
    return countMap
}
