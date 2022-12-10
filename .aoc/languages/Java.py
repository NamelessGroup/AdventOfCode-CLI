from .Language import Language
from typing import Tuple

class Java(Language):

    def getFormattedName(self) -> str:
        return "Java"

    def getTemplateDirectory(self) -> str:
        return ".aoc/file_templates/java/"

    def postCopy(self, day: int, target: str) -> None:
        pass

    def getPreRunCommand(self, day: str, task: int, sourcePath: str) -> [str | Tuple[str, str]]:
        return [f"javac {sourcePath}Runner.java {sourcePath}Task1.java {sourcePath}Task2.java"]

    def getRunCommand(self, day: str, task: int, sourcePath: str) -> str:
        return f"java {sourcePath}Runner {task}"

    def getTestCommand(self, day: int, task: int, sourcePath: str) -> str:
        return f"java {sourcePath}Runner {task} test"

    def hasIndividualTaskRunCommands(self) -> bool:
        return True

