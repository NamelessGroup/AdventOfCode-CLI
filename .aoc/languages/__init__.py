from .Python import Python
from .Java import Java

LANGUAGES_MAP = {
    "python": Python,
    "java": Java
}

LANGUAGES = LANGUAGES_MAP.keys()

def getLanguage(lang: str, langOptions) -> Language:
    return LANGUAGES_MAP[lang](langOptions)