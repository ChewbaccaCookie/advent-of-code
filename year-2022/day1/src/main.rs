use std::fs;

const TOP_NUM: usize = 3;

fn main() {
    println!(
        "Max Calorines: {}",
        load_file_and_find_max_calorines("./input.txt")[0]
    );
    println!(
        "Top Calorines summed up: {}",
        calc_top_calorines("./input.txt")
    );
}

pub fn load_file_and_find_max_calorines(file_name: &str) -> [u32; TOP_NUM] {
    let file_content = load_input(file_name);
    find_max_calorines(&file_content)
}

pub fn calc_top_calorines(file_name: &str) -> u32 {
    let calorines = load_file_and_find_max_calorines(file_name);
    let mut calculated = 0;
    for calorine in calorines {
        calculated += calorine;
    }
    calculated
}

fn find_max_calorines(file_content: &str) -> [u32; TOP_NUM] {
    let mut max_calorines = [0; TOP_NUM];
    let mut current_calorines = 0;

    for line in file_content.lines() {
        if line.is_empty() {
            for num in 0..TOP_NUM {
                if current_calorines > max_calorines[num] {
                    for arr_num in (num + 1..TOP_NUM).rev() {
                        max_calorines[arr_num] = max_calorines[arr_num - 1];
                    }
                    max_calorines[num] = current_calorines;
                    break;
                }
            }
            current_calorines = 0;
        } else {
            if let Ok(line_calorines) = u32::from_str_radix(line, 10) {
                current_calorines += line_calorines;
            }
        }
    }

    max_calorines
}

fn load_input(file_path: &str) -> String {
    let mut file_content =
        fs::read_to_string(file_path).expect(&format!("Failed to load file {}!", file_path));
    // add a empty line to the end
    file_content.push_str("\n\n");
    file_content
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_star1() {
        assert_eq!(
            load_file_and_find_max_calorines("./example-input.txt")[0],
            24000
        )
    }
    #[test]
    fn test_star1_top() {
        assert_eq!(
            load_file_and_find_max_calorines("./example-input.txt"),
            [24000, 11000, 10000]
        )
    }
    #[test]
    fn test_star2() {
        assert_eq!(calc_top_calorines("./example-input.txt"), 45000)
    }
}
