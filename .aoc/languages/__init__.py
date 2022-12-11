from .Python import Python

LANGUAGES_MAP = {
    "python": Python
}

LANGUAGES = LANGUAGES_MAP.keys()

def getLanguage(lang: str, **langOptions) -> Language:
    return LANGUAGES_MAP[lang](langOptions)