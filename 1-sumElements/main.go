package main

import(
    "fmt"
)

func main() {
    input1 := []int{1,2,3,4,5,6,7,8,9,10}
    result := sumArray(input1)

    fmt.Printf("The response is: %d\n", result)
}

func sumArray(arr []int) int {
    sum := 0

    for _, x := range arr {
        sum += x
    }

    return sum
}
