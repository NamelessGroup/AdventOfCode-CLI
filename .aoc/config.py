import argparse
from datetime import date
import json
import languages

def getArgumentParser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(
        prog = "Advent of Code-CLI",
        description = "Commandline-Interface for interacting with the Advent of Code more easily",
    )
    parser.add_argument('-l', '--language', help="Select the programming language to be used.", choices=languages.LANGUAGES)
    subparsers = parser.add_subparsers(help="Command to execute", required=True, dest="command")

    init = subparsers.add_parser('init', help="Creates a new directory structure for a day")
    init.add_argument('day', type=int, nargs="?", default=date.today().day, help="Specify the day to create. Defaults to today.")

    run = subparsers.add_parser('run', help="Runs the code for a given day")
    run.add_argument('day', type=int, help="Specify the day to run.")

    test = subparsers.add_parser('test', help="Tests the code for a given day")
    test.add_argument('day', type=int, help="Specify the day to test.")

    return parser

def getFileConfig() -> dict:
    try:
        f = open("./.aoc/config.json")
    except OSError:
        # Config doesn't exist
        return {}
    jsonConfig = json.loads(f.read())
    return jsonConfig

def getConfig() -> dict:
    parser = getArgumentParser()
    result = vars(parser.parse_args())

    config = getFileConfig()
    config.update({k: v for k, v in result.items() if v is not None})

    return config

def validateConfig(config: dict) -> None:
    if config['language'] not in languages.LANGUAGES:
        raise ValueError(f"Language '{config['language']}' doesn't exist!")
    if config['day'] <= 0 or config['day'] >= 26:
        raise ValueError(f"Invalid day. Expected 1-25, was {config['day']}")