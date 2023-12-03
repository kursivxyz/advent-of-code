fn main(){
    let input = include_str!("./input/01.txt");
    println!("Part 1: {}", part1(input)); 
}

fn part1(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            (line.chars().find_map(|a| a.to_digit(10)).unwrap()) * 10 + 
            (line.chars().rev().find_map(|a| a.to_digit(10)).unwrap())
        })
        .sum()
}

