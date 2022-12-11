from .Language import Language
from typing import Tuple

class Java(Language):

    def getFormattedName(self) -> str:
        return "Java"

    def getTemplateDirectory(self) -> str:
        return ".aoc/file_templates/java/"

    def postCopy(self, day: int, target: str) -> None:
        packageName = f"day{str(day).rjust(2, '0')}"
        # replace package declaration
        with open(f"{target}/Task1.java", 'r') as file:
            content = file.read()
        content = content.replace('INSERT_PACKAGE_NAME_HERE', packageName)
        with open(f"{target}/Task1.java", 'w') as file:
            file.write(content)

        with open(f"{target}/Task2.java", 'r') as file:
            content = file.read()
        content = content.replace('INSERT_PACKAGE_NAME_HERE', packageName)
        with open(f"{target}/Task2.java", 'w') as file:
            file.write(content)

        with open(f"{target}/Runner.java", 'r') as file:
            content = file.read()
        content = content.replace('INSERT_PACKAGE_NAME_HERE', packageName)
        with open(f"{target}/Runner.java", 'w') as file:
            file.write(content)

    def getPreRunCommand(self, day: str, task: int, sourcePath: str) -> [str | Tuple[str, str]]:
        sourcePath = sourcePath[2::]
        return [f"javac {sourcePath}/Runner.java {sourcePath}/Task1.java {sourcePath}/Task2.java"]

    def getRunCommand(self, day: str, task: int, sourcePath: str) -> str:
        sourcePath = sourcePath[2::]
        return f"java {sourcePath}/Runner {task}"

    def getTestCommand(self, day: int, task: int, sourcePath: str) -> str:
        sourcePath = sourcePath[2::]
        return f"java {sourcePath}/Runner {task} test"

    def hasIndividualTaskRunCommands(self) -> bool:
        return True

