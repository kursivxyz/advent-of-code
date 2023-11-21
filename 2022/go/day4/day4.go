package main

import (
	"fmt"
	"kursiv/aoc/util"
	"strconv"
	"strings"
)


func main() {
    input := util.Get_input_str("input.txt")
    ans1 := part1(input)
    fmt.Println("Answer part1: ", ans1)
}


func part1(input []string) int {
    var sum int;
    for _, line := range input {
        separatorPos := strings.Index(line, ",")
        first := line[:separatorPos]
        second:= strings.TrimLeft(line[separatorPos:], ",")

        start1, end1 := getRanges(first)
        start2, end2 := getRanges(second)
    
        if (start1 >= start2 && end1 <= end2 || start2 >= start1 && end2 <= end1) {
            sum++
        }
    }
    return sum;
}


func getRanges(string string) (int, int) {
    parts := strings.Split(string, "-")
    start,_ := strconv.Atoi(parts[0])
    end,_ := strconv.Atoi(parts[1])
    return start, end
}
