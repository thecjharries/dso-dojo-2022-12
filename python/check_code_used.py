from random import randint, seed


def check_code_used(code: str) -> str:
    seed(code)
    if 1 == randint(0, 1):
        return code.lower()
    return ''
