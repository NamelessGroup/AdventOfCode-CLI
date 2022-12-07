from .Language import Language

class Python(Language):

    def getFormattedName(self):
        return "[dodger_blue1]Python[/dodger_blue1]"

    def getTemplateDirectory(self):
        return ".aoc/file_templates/python/"

    def getRunCommand(self, day, task, sourcePath):
        return f"python -u {sourcePath}/runner.py {str(task)}"

    def getTestCommand(self, day, task, sourcePath):
        return f"python -u {sourcePath}/runner.py {str(task)} test"

    def hasIndividualTaskRunCommands(self):
        return True