package main

import (
    "fmt"
    "testing"
)

func TestSumArray(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {[]int{1, 2, 3, 4, 5}, 15},
        {[]int{10, 20, 30}, 60},
        {[]int{1, -1, 2, -2}, 0},
        {[]int{}, 0},
    }

    for _, testCase := range tests {
        t.Run(fmt.Sprintf("sum of %v", testCase.input), func(t *testing.T) {
            response := sumArray(testCase.input)
            if response != testCase.expected {
                t.Errorf("sumArray(%v) = %d; want %d", testCase.input, response, testCase.expected)
            }
        })
    }
}

