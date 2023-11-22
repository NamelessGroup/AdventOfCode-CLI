import os
from task import task1, task2
import sys

def parse_input(input_file_path: str) -> list[str]:
    file = open(input_file_path)
    raw_lines = file.readlines()
    file.close()
    lines = []
    for line in raw_lines:
        lines.append(line.rstrip("\n"))
    return lines

if __name__ == "__main__":
    if len(sys.argv) <= 2 or sys.argv[2] == "main":
        input_file = os.path.dirname(__file__) + "/input.in"
    elif sys.argv[2] == "test":
        input_file = os.path.dirname(__file__) + "/test.in"
        if not os.path.exists(input_file):
            print("Test file doesn't exist!")
            exit(1)
    lines = parse_input(input_file)

    if len(sys.argv) <= 1 or sys.argv[1] == "1":
        task1(lines)
    if len(sys.argv) <= 1 or sys.argv[1] == "2":
        task2(lines)