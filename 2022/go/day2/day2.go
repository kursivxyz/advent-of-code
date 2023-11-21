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
    
    ans2 := part2(input)
    fmt.Println("Answer part 2: ", ans2)
}

func part1(input []string) int {
    var score int; 
    for _, line := range input {
        parts := strings.Split(line, " ")
        player := normalizePlayCode(parts[1])
        opponent := normalizePlayCode(parts[0])
        score += played(player) + result(opponent, player)
    } 
    return score
}

func part2(input []string) int {
    var score int;
    for _, line := range input {
        parts := strings.Split(line, " ")
        opponent := normalizePlayCode(parts[0])
        desiredOutcome := normalizeDesiredOutcome(parts[1])
        neededPlay := neededPlay(opponent, desiredOutcome)
        score += played(neededPlay) + result(opponent, neededPlay)
    } 
    return score;
}

func neededPlay(opponent string, desiredOutcome string) string {
    if (desiredOutcome == "WIN") {
        return getWinningPlay(opponent)
    }
    if (desiredOutcome == "LOSE") {
        return getLosingPlay(opponent)
    }
    return opponent
}

func getWinningPlay(opponent string) string {  
    if (opponent == "SCISSOR") {
        return "ROCK"
    } 
    if (opponent == "ROCK") {
        return "PAPER"
    }
    if (opponent == "PAPER") {
        return "SCISSOR"
    }
    return ""
}

func getLosingPlay(opponent string) string {
    if (opponent == "SCISSOR") {
        return "PAPER"
    }
    if (opponent == "PAPER") {
        return "ROCK"
    }
    if (opponent == "ROCK") {
        return "SCISSOR"
    }
    return ""
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

func normalizeDesiredOutcome(outcome string) string {
    if (outcome == "Z") {
        return "WIN"
    }
    if (outcome == "Y") {
        return "DRAW"
    }
    if (outcome == "X") {
        return "LOSE"
    }
    return "";
}
