from .Language import Language
from typing import Tuple


class Java(Language):

    def getFormattedName(self) -> str:
        return "Java"

    def getTemplateDirectory(self) -> str:
        return ".aoc/file_templates/java/"

    def getPreRunCommand(self, day: str, task: int, sourcePath: str) -> [str | Tuple[str, str]]:
        return [f"javac Runner.java Tasks.java"]

    def getRunCommand(self, day: str, task: int, sourcePath: str) -> str:
        return f"java Runner {task} {sourcePath}\\"

    def getTestCommand(self, day: int, task: int, sourcePath: str) -> str:
        return f"java Runner {task} {sourcePath}\\ test"

    def hasIndividualTaskRunCommands(self) -> bool:
        return True

    def getRunCwd(self, day: str, task: int, sourcePath: str) -> str | None:
        return sourcePath

    def getPreRunCwd(self, day: str, task: int, sourcePath: str) -> str | None:
        return sourcePath
