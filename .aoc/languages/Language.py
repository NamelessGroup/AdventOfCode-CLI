from abc import ABC, abstractmethod

class Language(ABC):

    @abstractmethod
    def getFormattedName(self) -> str:
        pass

    @abstractmethod
    def getTemplateDirectory(self) -> str:
        pass

    def postCopy(self, day: int, target: str) -> None:
        pass

    @abstractmethod
    def getRunCommand(self, day: str, task: int, sourcePath: str) -> [str]:
        pass

    @abstractmethod
    def getTestCommand(self, day: int, task: int, sourcePath: str) -> [str]:
        pass

    @abstractmethod
    def hasIndividualTaskRunCommands(self) -> bool:
        pass