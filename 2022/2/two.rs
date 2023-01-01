use std::io;
use std::time::Instant;

fn main() {
    let began = Instant::now();
    let score = io::stdin()
        .lines()
        .filter_map(|line| line.ok())
        .fold(0, |score, round| {
            score
                + match &round[2..] {
                    "Y" => 3,
                    "Z" => 6,
                    _ => 0,
                }
                + match &round[..] {
                    "C Z" | "A Y" | "B X" => 1,
                    "A Z" | "B Y" | "C X" => 2,
                    "B Z" | "C Y" | "A X" => 3,
                    _ => 0,
                }
        });

    println!("{} took {:?}", score, began.elapsed());
}
