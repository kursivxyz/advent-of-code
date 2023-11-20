package main

import (
	"fmt"
	"kursiv/aoc/util"
)

func main() {
    input := util.Get_input_int("input.txt")
    ans1 := part1(input) 
    fmt.Println("Anwer part 1: ", ans1)

    
    ans2 := part2(input) 
    fmt.Println("Anwer part 2: ", ans2)
}

func part1(input []int) int {
    total := []int{}
    var sum int;
    for _, value := range input { 
        if value == 0 && sum != 0 {
            total = append(total, sum) 
            sum = 0
        }
        sum += value 
    }
    ans := util.FindMax(total)
    return ans
}

func part2(input []int) int {
    total := []int{}
    var sum int;
    for _, value := range input { 
        if value == 0 {
            total = append(total, sum) 
            sum = 0
        }
        sum += value 
    }
    ans := util.FindTopX(total, 3)

    return ans
}
