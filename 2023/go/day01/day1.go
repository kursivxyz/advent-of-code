package main

import (
	"fmt"
	"kursiv/aoc/util"
	"regexp"
	"strconv"
)

func main() {
    input := util.Get_input_str("input.txt")
    ans1 := part1(input)
    fmt.Println("Answer part1: ", ans1)
}

func part1(input []string) int {
    var sum int; 

    regex, err := regexp.Compile("[0-9]")
    if err != nil {
        panic(err)
    }

    for _,v := range input {
        ints := regex.FindAllString(v, -1) 
        digits := ints[0] + ints[len(ints)-1]
        number, err := strconv.Atoi(digits) 
        if err != nil {
            panic(err)
        }
        sum += number
    }
    return sum 
}
