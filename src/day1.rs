use std::num::ParseIntError;

use aoc_runner_derive::{aoc, aoc_generator};

#[aoc_generator(day1)]
pub fn input_generator(input: &str) -> Result<Vec<u16>, ParseIntError> {
    input.lines().map(|l| l.parse::<u16>()).collect()
}

#[aoc(day1, part1)]
pub fn solve_part1(input: &Vec<u16>) -> usize {
    input
        .windows(2)
        .map(|x| match x {
            [x, y] => y > x,
            &_ => unreachable!(),
        })
        .filter(|&d| d == true)
        .count()
}

#[aoc(day1, part2)]
pub fn solve_part2(input: &Vec<u16>) -> usize {
    let depths: Vec<u16> = input
        .windows(3)
        .map(|x| match x {
            [x, y, z] => x + y + z,
            &_ => unreachable!(),
        })
        .collect();

    solve_part1(&depths)
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_input_generator() {
        let input = "5\n7\n49\n";
        assert_eq!(input_generator(input).unwrap(), vec![5, 7, 49]);
    }

    #[test]
    fn test_solve_part1() {
        let input = vec![199, 200, 208, 210, 200, 207, 240, 269, 260, 263];
        assert_eq!(solve_part1(&input), 7);
    }

    #[test]
    fn test_solve_part2() {
        let input = vec![199, 200, 208, 210, 200, 207, 240, 269, 260, 263];
        assert_eq!(solve_part2(&input), 5);
    }
}
