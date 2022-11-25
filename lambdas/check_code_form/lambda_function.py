from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.typing import LambdaContext
from re import compile, IGNORECASE
from typing import TypedDict

CODE_FORM_PATTERN = compile(r'^[a-z]{5}-\d{5}-[a-z]{5}$', IGNORECASE)
logger = Logger()


class InputEvent(TypedDict):
    code: str


class OutputEvent(TypedDict):
    success: bool
    code: str


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = event['code']
    if CODE_FORM_PATTERN.match(code):
        return {
            'success': True,
            'code': code.lower(),
        }
    return {
        'success': False,
        'code': '',
    }
