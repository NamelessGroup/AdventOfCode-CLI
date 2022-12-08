# Advent of Code CLI

> TODO: Write better documentation

A commandline interface to interact with the [Advent of Code](https://adventofcode.com).

## Usage:
- Install it by cloning the repo, or using this repo as a template
- Access the CLI using the `aoc` or `aoc.ps1` scripts.
## Commands / Features:
- `aoc init [day]` - Initializes a folder for the day, copies the template files of the given language and downloads your puzzle input.
- `aoc run <day>` - Runs the code for the supplied day with the puzzle input. Also measures execution time.
- `aoc test <day>` - Runs the code for the supplied day with testing input. Also measures execution time.

## Supported Languages:
At the moment, only Python, however, Java (and some others) are Work-In-Progress. (Maybe even by you - open a PR ;) )

## Contributing
> TODO
### Adding new Languages
- All languages are a python file in the `.aoc/languages` directory.
- Add a new file with the language name, create a class in it, and extend the `Language` class.
- Register the language in the `.aoc/languages/__init__.py` file.
- (Take a look at other, existing languages for inspiration and reference)