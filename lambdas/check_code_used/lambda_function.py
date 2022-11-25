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


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = event['code']
    seed(code)
    if 1 == randint(0, 1):
        return {
            'success': True,
            'code': code.lower(),
        }
    return {
        'success': False,
        'code': '',
    }
