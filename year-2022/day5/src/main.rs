use regex::Regex;

fn main() {
    get_top_crates("./input.txt");
}

fn get_top_crates(file_path: &str) -> &str {
    let file_content = shared::load_input(file_path);
    let (width, height) = find_max_x(&file_content);
    let crates = parse_crates(width, height, &file_content);
    move_crates(crates, width, height, &file_content);
    println!("{}, {}", width, height);
    ""
}

fn find_max_x(file_content: &str) -> (u32, u32) {
    let mut height: usize = 0;
    for line in file_content.lines() {
        if line.is_empty() {
            break;
        }
        height += 1;
    }
    let last_line = file_content.lines().nth(height - 1);
    let last_line_chars = last_line.unwrap().chars();
    let width = last_line_chars.rev().nth(1).unwrap();

    (width.to_digit(10).unwrap(), (height - 1) as u32)
}

fn parse_crates(width: u32, height: u32, file_content: &str) -> Vec<char> {
    let mut crates = vec![' '; (width * height) as usize];

    let mut line_num = file_content.lines().count() as u32;
    let mut y = 0;
    for line in file_content.lines().rev() {
        if line_num > height {
            line_num -= 1;
            continue;
        }
        let mut x = 0;
        let mut i = 0;
        for c in line.chars() {
            if i % 4 == 1 {
                let index = (width * y + x) as usize;
                crates[index] = c;
                x += 1;
            }
            i += 1;
        }
        y += 1;
        line_num -= 1;
    }
    crates
}

fn move_crates(crates: Vec<char>, width: u32, height: u32, file_content: &str) {
    let reg = Regex::new(r"move (\d+) from (\d+) to (\d+)");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(get_top_crates("./example-input.txt"), "CMZ")
    }
}
