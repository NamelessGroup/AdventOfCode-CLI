import os
from task import task1, task2
import sys

def parse_input(input_file_path):
    file = open(input_file_path)
    raw_lines = file.readlines()
    file.close()
    lines = []
    for line in raw_lines:
        lines.append(line.rstrip("\n"))
    return lines

if __name__ == "__main__":
    if len(sys.argv) <= 2 or sys.argv[2] == "main":
        input_file = os.path.dirname(__file__) + "/input"
    elif sys.argv[2] == "test":
        input_file = os.path.dirname(__file__) + "/test"
        if not os.path.exists(input_file):
            print("Test file doesn't exist!")
            exit(1)
    lines = parse_input(input_file)

    if len(sys.argv) <= 1 or sys.argv[1] == "1":
        try:
            task1(lines)
        except Exception as e:
            print(e)
            exit(1)
    if len(sys.argv) <= 1 or sys.argv[1] == "2":
        try:
            task2(lines)
        except Exception as e:
            print(e)
            exit(1)