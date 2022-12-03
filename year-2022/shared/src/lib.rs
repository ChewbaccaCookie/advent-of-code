use std::fs;

pub fn load_input(file_path: &str) -> String {
    let file_content =
        fs::read_to_string(file_path).expect(&format!("Failed to load file {}!", file_path));
    file_content
}
