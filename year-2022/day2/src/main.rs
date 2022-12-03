use regex::Regex;
use std::error::Error;

fn main() {
    calculate_score("./input.txt");
}

fn calculate_score(file_path: &str) -> Result<u32, Box<dyn Error>> {
    let file_content = shared::load_input(file_path);
    let file_re = Regex::new(r"([ABC])\s([XYZ])").unwrap();
    let mut score = 0;

    for line in file_re.captures_iter(&file_content) {
        let left = line.get(1).unwrap().as_str();
        let right = line.get(2)?.as_str();
    }

    score
}

mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(calculate_score("./example-input.txt").unwrap(), 15)
    }
}
