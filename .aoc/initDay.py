import os
import requests
import shutil

AOC_BASE_URL = "https://adventofcode.com"
AOC_YEAR = 2022

def makeNewDay(day, console, lang):
    newDirectoryPath = f"src/day{str(day).rjust(2, '0')}"

    # Create new directory
    if not os.path.exists(newDirectoryPath):
        os.mkdir(newDirectoryPath)
    console.log(f"Created directory {newDirectoryPath}.")

    # Copy files
    shutil.copytree(lang.getTemplateDirectory(), newDirectoryPath + "/", dirs_exist_ok=True)
    console.log(f"Copied {lang.getFormattedName()} template.")

def getInput(day, console):
    if not os.path.exists(".aoc/SESSION_TOKEN"):
        console.log("[red]You need to write your session token to the file '.aoc/SESSION_TOKEN' in order to download your input!")
        return
    tokenFile = open(".aoc/SESSION_TOKEN")
    AOC_SESSION_COOKIE = tokenFile.read()
    tokenFile.close()

    response = requests.get(url=f"{AOC_BASE_URL}/{AOC_YEAR}/day/{str(day)}/input", cookies={"session": AOC_SESSION_COOKIE})
    if response.ok:
        data = response.text
        file = open(f"src/day{str(day).rjust(2, '0')}/input", "w+")
        file.write(data.rstrip("\n"))
        file.close()
        console.log("Successfully downloaded input file")
    else:
        console.log("[red]Error while getting input")