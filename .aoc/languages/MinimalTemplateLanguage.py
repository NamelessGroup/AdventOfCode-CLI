# Use this language file as a template
# Register it in `languages/__init__.py`
from .Language import Language

class TemplateLanguage(Language):

    def getFormattedName(self):
        # Return the name of the programming language, with optional markup for the rich-library
        return "[red]TemplateLanguage[/red]"
    
    def getTemplateDirectory(self):
        # Return the path to the template directory to copy
        return ".aoc/file_templates/template"

    def getRunCommand(self, day, task, sourcePath):
        # Return the command to execute the specified day & task
        # May execute both tasks a day
        return "echo 'Heyo'"
    
    def getTestCommand(self, day, task, sourcePath):
        # Return the command to test the specified day & task
        # May test both tasks a day
        return "echo 'Test'"
    
    def hasIndividualTaskRunCommands(self):
        # Return whether the getRunCommand & getTestCommand executes both tasks at once (False)
        # of if they only execute the task given to them (True).
        return True