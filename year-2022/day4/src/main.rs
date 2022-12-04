use regex::Regex;

fn main() {
    let (over_pairs, between_pairs) = find_assignment_pairs("./input.txt");
    println!("Star 1: {}\nStart 2: {}", over_pairs, between_pairs);
}

fn find_assignment_pairs(file_path: &str) -> (u32, u32) {
    let mut over_pairs = 0;
    let mut between_pairs = 0;
    let regex = Regex::new(r"^(\d+)-(\d+),(\d+)-(\d+)$").unwrap();
    let file_content = shared::load_input(file_path);

    for line in file_content.lines() {
        let captures = regex.captures(line).unwrap();
        let left_start = shared::get_capture_value_as_int(&captures, 1);
        let left_end = shared::get_capture_value_as_int(&captures, 2);
        let right_start = shared::get_capture_value_as_int(&captures, 3);
        let right_end = shared::get_capture_value_as_int(&captures, 4);

        if (left_start >= right_start && left_end <= right_end)
            || (left_start <= right_start && left_end >= right_end)
        {
            over_pairs += 1;
        }
        if left_start <= right_end && left_end >= right_start {
            between_pairs += 1;
        }
    }

    (over_pairs, between_pairs)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(find_assignment_pairs("./example-input.txt").0, 2)
    }

    #[test]
    fn test_star2() {
        assert_eq!(find_assignment_pairs("./example-input.txt").1, 4)
    }
}
