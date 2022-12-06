import sys
from datetime import date
from rich.progress import Progress
from rich import print
from rich.console import Console
from runner import runDay, testDay
from initDay import getInput, makeNewDay
from config import getConfig, validateConfig
import languages

if __name__ == '__main__':
    console = Console()

    config = getConfig()
    try:
        validateConfig(config)
    except ValueError as e:
       console.print(f"[red]{e}")
       exit(1)

    lang = languages.getLanguage(config['language'])
    day = config['day']
    
    if config['command'] == "init":
        with console.status(f"Initialising day {day}..."):
            makeNewDay(day, console, lang)
            getInput(day, console)

    elif config['command'] == "run":
        with console.status(f"Running day {day}..."):
            runDay(day, console, lang)
    elif config['command'] == "test":
        with console.status(f"Running day {day}..."):
            testDay(day, console, lang)
    else:
        console.print(f"[red]Invalid command '{config['command']}'")