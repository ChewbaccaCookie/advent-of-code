use std::fs;

use regex::Captures;

pub fn load_input(file_path: &str) -> String {
    let file_content =
        fs::read_to_string(file_path).expect(&format!("Failed to load file {}!", file_path));
    file_content
}

pub fn get_capture_value_as_int(captures: &Captures, index: usize) -> u32 {
    captures
        .get(index)
        .unwrap()
        .as_str()
        .parse::<u32>()
        .unwrap()
}
