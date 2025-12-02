use std::env;
use std::fs;

mod task;

fn parse_input(input_file_path: &str) -> Vec<String> {
    let contents = fs::read_to_string(input_file_path)
        .expect("Failed to read input file");
    
    contents
        .lines()
        .map(|line| line.to_string())
        .collect()
}

fn main() {
    let args: Vec<String> = env::args().collect();
    
    let task_number = if args.len() > 1 {
        args[1].as_str()
    } else {
        "1"
    };
    
    let input_mode = if args.len() > 2 {
        args[2].as_str()
    } else {
        "main"
    };
    
    let input_file = if input_mode == "test" {
        "test.in"
    } else {
        "solve.in"
    };
    
    if input_mode == "test" && !std::path::Path::new(input_file).exists() {
        eprintln!("Test file doesn't exist!");
        std::process::exit(1);
    }
    
    let lines = parse_input(input_file);
    
    match task_number {
        "1" => task::task1(&lines),
        "2" => task::task2(&lines),
        _ => {
            task::task1(&lines);
            task::task2(&lines);
        }
    }
}
