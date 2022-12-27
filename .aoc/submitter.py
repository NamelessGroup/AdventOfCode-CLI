import os

import requests


def submit(year, day, task, console, answer):
    if not (day == 1 or day == 2):
        console.log("[red]Wrong task given")
        return False

    if not os.path.exists(".aoc/SESSION_TOKEN"):
        console.log("[red]You need to write your session token to the file '.aoc/SESSION_TOKEN' in order to download your input!")
        return False
    tokenFile = open(".aoc/SESSION_TOKEN")
    AOC_SESSION_COOKIE = tokenFile.read()
    tokenFile.close()

    url = f"https://adventofcode.com/{year}/day/{day}/answer"
    post = {
        'level': task,
        'answer': answer
    }
    result = requests.post(url=url, cookies={"session": AOC_SESSION_COOKIE}, data=post)
    if not result.ok:
        console.log("[red]Error while submitting")
        return False
    if result.text.__contains__("That's not the right answer"):
        console.log("[red]That was not the right answer")
    elif result.text.__contains__("You don't seem to be solving the right level"):
        console.log("[green]You already completed this puzzle")
    elif result.text.__contains__("That's the right answer"):
        console.log("[green]That's the right answer")
        return True
    elif result.text.__contains__("You gave an answer too recently"):
        console.log("[red]You gave an answer too recently. Please wait a moment")
    else:
        console.log("[red]Error while submitting")
        console.log(result.text)
    return False
