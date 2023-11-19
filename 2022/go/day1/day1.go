package main

import (
	"fmt"
	"kursiv/aoc/util"
	"sort"
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
        if value == 0 {
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
    sort.Ints(total)
    var topSum int
    for i := 0; i < 3; i++ {
        topSum += total[len(total)-1-i]
    }
    return topSum
}
