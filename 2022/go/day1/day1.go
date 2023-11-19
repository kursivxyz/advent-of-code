package main

import (
	"fmt"
    "kursiv/aoc/util"
)

func main() {
    lines := util.Get_input_int("input.txt")
    total := []int{}
    var sum int;
    for _, value := range lines { 
        if value == 0 {
            total = append(total, sum) 
            sum = 0
        }
        sum += value 
    }
    ans := util.FindMax(total)
    fmt.Println("Anwer: ", ans)
}
