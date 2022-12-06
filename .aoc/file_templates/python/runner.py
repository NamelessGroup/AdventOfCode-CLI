import os
from rich.console import Console
from task import task1, task2
import sys

def parseInput(inputFilePath):
    file = open(inputFilePath)
    rawLines = file.readlines()
    file.close()
    lines = []
    for line in rawLines:
        lines.append(line.rstrip("\n"))
    return lines

if __name__ == "__main__":
    console = Console()

    if len(sys.argv) <= 2 or sys.argv[2] == "main":
        inputFile = os.path.dirname(__file__) + "/input"
    elif sys.argv[2] == "test":
        inputFile = os.path.dirname(__file__) + "/test"
        if not os.path.exists(inputFile):
            console.print("Test file doesn't exist!")
            exit(1)
    lines = parseInput(inputFile)

    if len(sys.argv) <= 1 or sys.argv[1] == "1":
        try:
            task1(lines)
        except Exception:
            console.print_exception()
            exit(1)
    if len(sys.argv) <= 1 or sys.argv[1] == "2":
        try:
            task2(lines)
        except Exception:
            console.print_exception()
            exit(1)