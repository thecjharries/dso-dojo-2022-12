from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.typing import LambdaContext
from typing import TypedDict

logger = Logger()
PROTOCOL = 'https'
ROOT_URL = 'reallycoolwebsite.com'


class InputEvent(TypedDict):
    code: str


class OutputEvent(TypedDict):
    success: bool
    code: str
    link: str


def generate_link(code: str) -> str:
    return f'{PROTOCOL}://{ROOT_URL}/download/{code}'


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = event['code']
    return {
        'success': True,
        'code': code,
        'link': generate_link(code),
    }
