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

    ans2 := part2(input)
    fmt.Println("Answer part2: ", ans2)
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

func part2(input []string) int {
    var sum int;

    for i := 0; i < len(input); i += 3 {
        first, second, third := input[i], input[i+1], input[i+2] 

        sharedChars := make(map[rune]bool)
        for _, item := range first {
            if strings.ContainsRune(second, item) && strings.ContainsRune(third, item) {
                if (sharedChars[item]) {
                    break
                }
                sharedChars[item] = true
                sum += getPriority(string(item))
            }
         }
    }
    return sum
} 


func getPriority(letter string) int {
    return strings.Index(alphabet, letter) + 1 
}
