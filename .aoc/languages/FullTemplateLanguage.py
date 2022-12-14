# Use this language file as a template
# Register it in `languages/__init__.py`
from .Language import Language

class TemplateLanguage(Language):

    def __init__(self, options):
        # Optional method
        # Options is a dict with language-specific options
        pass

    def getFormattedName(self):
        # Return the name of the programming language, with optional markup for the rich-library
        return "[red]TemplateLanguage[/red]"
    
    def getTemplateDirectory(self):
        # Return the path to the template directory to copy
        return ".aoc/file_templates/template"
    
    def postCopy(self, day, target):
        # Optional method
        # Is ran after the template has been copied to path, if e.g. replacements in the files need to be made
        pass

    def getPreRunCommand(self, day, task, sourcePath):
        # Optional method
        # Return commands to be ran before executing a task (e.g. build scripts)
        # Can either be a list of commands, or a list of tuples, containing ("command_label", "command") where label is displayed in the CLI instead of the full command
        return []

    def getPreRunCwd(self, day, task, sourcePath):
        # Optional method
        # Return the cwd for the preRunCommands
        return None

    def getRunCommand(self, day, task, sourcePath):
        # Return the command to execute the specified day & task
        # May execute both tasks a day
        return "echo 'Heyo'"
    
    def getRunCwd(self, day, task, sourcePath):
        # Optional method
        # Return the current working directory for the run command
        return None
    
    def getTestCommand(self, day, task, sourcePath):
        # Return the command to test the specified day & task
        # May test both tasks a day
        return "echo 'Test'"
    
    def hasIndividualTaskRunCommands(self):
        # Return whether the getRunCommand & getTestCommand executes both tasks at once (False)
        # of if they only execute the task given to them (True).
        return True