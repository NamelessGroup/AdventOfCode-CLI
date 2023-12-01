# Advent of Code CLI

A commandline interface to interact with the [Advent of Code](https://adventofcode.com).

## Usage:
- Install it by downloading the latest release or building from scratch.

## Commands / Features:
#### `aoc-cli init [-l/--lang language] [-d/--day day] [-y/--year year] [--debug] [--no-emojis] [-c/--cookie cookie] [--task2]`
Initializes a folder for the day, copies the template files of the given language and downloads your puzzle input.
By default, uses the current day.
If `--task2` is given, also tries to download the second challenge description.

#### `aoc-cli solve [task] [-l/--lang language] [-d/--day day] [-y/--year year] [--debug] [--no-emojis] [-c/--cookie cookie] [--submit]`
Runs the code for the supplied day with the puzzle input. Also measures execution time.
By default, uses the current day and task 1.
If `--submit` is given, also tries to submit your solution, which is the last line of the output of your script, as the solution to the website.

#### `aoc-cli test [task] [-l/--lang language] [-d/--day day] [-y/--year year] [--debug] [--no-emojis]`
Runs the code for the supplied day with testing input. Also measures execution time.
By default, uses the current day and task 1.

#### `aoc-cli config list|[key] [value] [-l/--lang language]`
Updates the config file.
If `list` is given, lists all possible configuration values, including language specific configuration.
See below for configuration file reference.

## Configuration
Persistent configuration is saved inside the `aoc-cli-config.json` file, and can be edited via a text editor, or the `aoc-cli config` command.

### Configuration keys

| Key        | Equivalent cli flag | Description                                          |
| ---------- | ------------------- | ---------------------------------------------------- |
| `language` | `-l/--language`     | Sets the default language.                           |
| `noEmojis` | `--no-emojis`       | When set, disables emojis in the output.             |
| `cookie`   | `-c/-cookie`        | Sets the cookie for authenticating with the website. |

### Language specific configuration:

#### Python

| Key                           | Default  | Description               |
| ----------------------------- | -------- | ------------------------- |
| `languages.python.executable` | `python` | Python executable to run. |


## Supported Languages:
- Python
- TypeScript
- Java
- Haskell

## Contributing

### Adding new Languages
- All languages are a go file in the `runner/languages/` directory.
- All language template files are embedded in the executable and inside the `runner/languages/[LANGUAGE]` directory.
- Add a new file with the language name, create a class in it, and implement the `Language` interface.
- Register the language in the `runner/languages.go` file.
- (Take a look at other, existing languages for inspiration and reference)