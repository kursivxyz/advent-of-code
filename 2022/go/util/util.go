package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Get_input_str(input string) []string {
    content, err := os.ReadFile(input)
    if err != nil {
        log.Fatal("something gon wrong lol")
    }

    return strings.Split(string(content), "\n")
}

func Get_input_int(input string) []int {
    content, err := os.ReadFile(input)
    if err != nil {
        log.Fatal("something gon wrong lol")
    }

    lines := strings.Split(string(content), "\n")
    var ints []int
    for _, line := range lines {
        if line == "" {
            line = "0"
        }
        n, err := strconv.Atoi(line) 
        if err != nil {
            log.Fatalf("error during conversion of %s", err)
        }
        ints = append(ints, n)
    }
    return ints;
}

func FindMax(ints []int) int {
    max := ints[0] 
    for _, value := range ints {
        if value > max {
            max = value
        }
    } 
    return max
}
