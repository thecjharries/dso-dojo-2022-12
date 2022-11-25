from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.typing import LambdaContext
from random import randint, seed
from typing import TypedDict

logger = Logger()


class InputEvent(TypedDict):
    code: str


class OutputEvent(TypedDict):
    success: bool
    code: str

def check_code_used(code: str) -> str:
    seed(code)
    if 1 == randint(0, 1):
        return code.lower()
    return ''


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = check_code_used(event['code'])
    return {
        'success': '' != code,
        'code': code,
    }
