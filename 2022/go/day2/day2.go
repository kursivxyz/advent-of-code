package main

import (
    "fmt"
    "kursiv/aoc/util"
    "strings"
)

func main() {
    input := util.Get_input_str("input.txt")
    ans1 := part1(input)
    fmt.Println("Answer part 1: ", ans1)
}

func part1(input []string) int {
    fmt.Println(len(input))
    var score int; 
    for _, line := range input {
        parts := strings.Split(line, " ")
        player := normalizePlayCode(parts[1])
        opponent := normalizePlayCode(parts[0])
        score += played(player) + result(opponent, player)
    } 
    return score
}

func played(played string) int {
    if (played == "ROCK") {
        return 1
    } 
    if (played == "PAPER") {
        return 2
    } 
    if (played == "SCISSOR") {
        return 3
    }
    return 0
}

func result(opponent string, player string) int {
    if (opponent == "ROCK" && player == "PAPER" || opponent == "PAPER" && player == "SCISSOR" || opponent == "SCISSOR" && player == "ROCK") {
        return 6
    }
    if (opponent == player) { 
        return 3 
    }
    return 0
}

func normalizePlayCode(played string) string {
    if (played == "A" || played == "X")  {
        return "ROCK"
    } 
    if (played == "B" || played == "Y")  {
        return "PAPER"
    } 
    if (played == "C" || played == "Z")  {
        return "SCISSOR"
    }
    return "";
}
