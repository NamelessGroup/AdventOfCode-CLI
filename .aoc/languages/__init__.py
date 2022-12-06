from .Python import Python

LANGUAGES_MAP = {
    "python": Python
}

LANGUAGES = LANGUAGES_MAP.keys()

def getLanguage(lang: str) -> Language:
    return LANGUAGES_MAP[lang]()