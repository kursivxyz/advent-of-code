use std::collections::HashMap;

fn main(){
    let input = include_str!("./input/01.txt");
    println!("Part 1: {}", part1(input)); 
    println!("Part 2: {}", part2(input));
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

fn part2(input: &str) -> u32 {
        let valid_digits: HashMap<&str, u32>  = HashMap::from([
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
        ("1", 1),
        ("2", 2),
        ("3", 3),
        ("4", 4),
        ("5", 5),
        ("6", 6),
        ("7", 7),
        ("8", 8),
        ("9", 9),
        ("0", 0),]);

        input
            .lines()
            .map(|line| {
                valid_digits[(0..line.len()).find_map(|i| valid_digits.keys().find(|x| line[i..].starts_with(*x))).unwrap()] * 10 +
                valid_digits[(0..line.len()).find_map(|i| valid_digits.keys().find(|x| line[..line.len()-i].ends_with(*x))).unwrap()]
            })
            .sum()
}
