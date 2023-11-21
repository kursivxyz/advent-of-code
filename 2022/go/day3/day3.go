package main

import (
    "fmt"
    "kursiv/aoc/util"
    "strings"
)

func main() {
    input := util.Get_input_str("input.txt")
    ans1 := part1(input)
    fmt.Println("Answer part1: ", ans1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1(input []string) int {
    var sum int;
    for _, line := range input {
        mid := len(line)/2
        first := line[:mid]
        last := line[mid:]

        charsInFirst := make(map[rune]bool)
        for _, item := range first {
            charsInFirst[item] = true
        }

        for _, item := range last {
            if (charsInFirst[item]) {
                sum += getPriority(string(item))
                delete(charsInFirst, item)
            }
        }
    } 
    return sum;
}



func getPriority(letter string) int {
    return strings.Index(alphabet, letter) + 1 
}
