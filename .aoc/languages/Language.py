from abc import ABC, abstractmethod
from typing import Tuple

class Language(ABC):

    def __init__(self, **options):
        pass

    @abstractmethod
    def getFormattedName(self) -> str:
        pass

    @abstractmethod
    def getTemplateDirectory(self) -> str:
        pass

    def postCopy(self, day: int, target: str) -> None:
        pass

    def getPreRunCommand(self, day: str, task: int, sourcePath: str) -> [str | Tuple[str, str]]:
        return []

    @abstractmethod
    def getRunCommand(self, day: str, task: int, sourcePath: str) -> str:
        pass

    @abstractmethod
    def getTestCommand(self, day: int, task: int, sourcePath: str) -> str:
        pass

    @abstractmethod
    def hasIndividualTaskRunCommands(self) -> bool:
        pass