from re import compile, IGNORECASE

CODE_FORM_PATTERN = compile(r'^[a-z]{5}-\d{5}-[a-z]{5}$', IGNORECASE)


def check_code_form(code: str) -> str:
    if CODE_FORM_PATTERN.match(code):
        return code.lower()
    return ''
