use regex::Regex;

fn main() {
    let (score1, score2) = calculate_score("./input.txt");
    println!("Score Star 1: {score1} \nScore Start 2: {score2}",);
}

fn calculate_score(file_path: &str) -> (u32, u32) {
    let scores: Vec<Vec<u32>> = vec![vec![3, 0, 6], vec![6, 3, 0], vec![0, 6, 3]];
    let game_options: Vec<Vec<u32>> = vec![vec![2, 0, 1], vec![0, 1, 2], vec![1, 2, 0]];

    let reg: Regex = Regex::new(r"([ABC])\s([XYZ])").unwrap();

    let file_content = shared::load_input(file_path);
    let mut star1_score = 0;
    let mut star2_score = 0;

    for line in reg.captures_iter(&file_content) {
        let left = line.get(1).unwrap().as_str();
        let right = line.get(2).unwrap().as_str();

        let left_score = left.chars().next().unwrap() as u32 - 65;
        let mut right_score = right.chars().next().unwrap() as u32 - 88;

        star1_score += right_score
            + 1
            + scores
                .get(right_score as usize)
                .unwrap()
                .get(left_score as usize)
                .unwrap();

        right_score = *game_options
            .get(right_score as usize)
            .unwrap()
            .get(left_score as usize)
            .unwrap();

        star2_score += right_score
            + 1
            + scores
                .get(right_score as usize)
                .unwrap()
                .get(left_score as usize)
                .unwrap();
    }

    (star1_score, star2_score)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(calculate_score("./example-input.txt").0, 15)
    }
    #[test]
    fn test_star2() {
        assert_eq!(calculate_score("./example-input.txt").1, 12)
    }
}
