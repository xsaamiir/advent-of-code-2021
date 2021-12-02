use std::str::FromStr;

use aoc_runner_derive::{aoc, aoc_generator};

#[derive(PartialEq, Debug)]
pub enum Direction {
    Forward,
    Down,
    Up,
}

impl FromStr for Direction {
    type Err = ();

    fn from_str(input: &str) -> Result<Direction, Self::Err> {
        match input {
            "forward" => Ok(Direction::Forward),
            "down" => Ok(Direction::Down),
            "up" => Ok(Direction::Up),
            _ => Err(()),
        }
    }
}

#[derive(PartialEq, Debug)]
pub struct Step(Direction, i32);

#[aoc_generator(day2)]
pub fn input_generator(input: &str) -> Option<Vec<Step>> {
    input
        .lines()
        .map(|l| {
            l.split_once(' ').map(|(d, u)| {
                let direction = Direction::from_str(d).unwrap();
                let unit = u.parse::<i32>().unwrap();
                Step(direction, unit)
            })
        })
        .collect()
}

pub struct Position {
    horizontal: i32,
    depth: i32,
    aim: i32,
}

impl Position {
    fn multiply(&self) -> i32 {
        self.horizontal * self.depth
    }
}

#[aoc(day2, part1)]
pub fn solve_part1(input: &Vec<Step>) -> i32 {
    input
        .iter()
        .fold(
            Position {
                horizontal: 0,
                depth: 0,
                aim: 0,
            },
            |p, s| match s {
                Step(Direction::Forward, u) => Position {
                    horizontal: p.horizontal + u,
                    ..p
                },
                Step(Direction::Up, u) => Position {
                    depth: p.depth - u,
                    ..p
                },
                Step(Direction::Down, u) => Position {
                    depth: p.depth + u,
                    ..p
                },
            },
        )
        .multiply()
}

#[aoc(day2, part2)]
pub fn solve_part2(input: &Vec<Step>) -> i32 {
    input
        .iter()
        .fold(
            Position {
                horizontal: 0,
                depth: 0,
                aim: 0,
            },
            |p, s| match s {
                Step(Direction::Forward, u) => Position {
                    horizontal: p.horizontal + u,
                    depth: p.depth + (p.aim * u),
                    ..p
                },
                Step(Direction::Up, u) => Position {
                    aim: p.aim - u,
                    ..p
                },
                Step(Direction::Down, u) => Position {
                    aim: p.aim + u,
                    ..p
                },
            },
        )
        .multiply()
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_input_generator() {
        let input = "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2";
        assert_eq!(
            input_generator(input).unwrap(),
            vec![
                Step(Direction::Forward, 5),
                Step(Direction::Down, 5),
                Step(Direction::Forward, 8),
                Step(Direction::Up, 3),
                Step(Direction::Down, 8),
                Step(Direction::Forward, 2),
            ]
        );
    }

    #[test]
    fn test_solve_part1() {
        let input = vec![
            Step(Direction::Forward, 5),
            Step(Direction::Down, 5),
            Step(Direction::Forward, 8),
            Step(Direction::Up, 3),
            Step(Direction::Down, 8),
            Step(Direction::Forward, 2),
        ];
        assert_eq!(solve_part1(&input), 150);
    }

    #[test]
    fn test_solve_part2() {
        let input = vec![
            Step(Direction::Forward, 5),
            Step(Direction::Down, 5),
            Step(Direction::Forward, 8),
            Step(Direction::Up, 3),
            Step(Direction::Down, 8),
            Step(Direction::Forward, 2),
        ];
        assert_eq!(solve_part2(&input), 900);
    }
}
