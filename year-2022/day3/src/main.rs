fn main() {
    let score = calculate_single_backpack_priorities("./input.txt");
    println!("Star 1 Score: {} ", score);

    let score = calculate_three_backpack_priorities("./input.txt");
    println!("Star 2 Score: {} ", score);
}

fn calculate_single_backpack_priorities(file_path: &str) -> u32 {
    let file_content = shared::load_input(file_path);
    let mut priority = 0;
    for line in file_content.lines() {
        // split the string into two parts
        let (left, right) = line.split_at(line.len() / 2);
        // store left keys into array
        let mut characters: [bool; 255] = [false; 255];
        for char in left.chars() {
            characters[char as usize] = true;
        }
        let mut duplicate: Option<char> = None;
        // Check for duplicate
        for char in right.chars() {
            if characters[char as usize] == true {
                duplicate = Some(char);
                break;
            }
        }

        let char_val = if (duplicate.unwrap() as u32) < 97 {
            duplicate.unwrap() as u32 - 38
        } else {
            duplicate.unwrap() as u32 - 96
        };
        priority += char_val;
    }

    priority
}

fn calculate_three_backpack_priorities(file_path: &str) -> u32 {
    let file_content = shared::load_input(file_path);
    let mut three_chars: [u8; 255] = [0; 255];
    let mut i = 0;
    let mut priority = 0;
    for line in file_content.lines() {
        if i % 3 == 0 {
            three_chars = [0; 255];
        }
        let mut single_chars: [bool; 255] = [false; 255];

        for char in line.chars() {
            // prevent duplicate counts per row
            if single_chars[char as usize] {
                continue;
            }

            three_chars[char as usize] += 1;
            single_chars[char as usize] = true;

            if i % 3 == 2 && three_chars[char as usize] == 3 {
                let char_val = if (char as u32) < 97 {
                    char as u32 - 38
                } else {
                    char as u32 - 96
                };
                priority += char_val;

                break;
            }
        }
        i += 1;
    }

    priority
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(
            calculate_single_backpack_priorities("./example-input.txt"),
            157
        )
    }
    #[test]
    fn test_star2() {
        assert_eq!(
            calculate_three_backpack_priorities("./example-input.txt"),
            70
        )
    }
}
