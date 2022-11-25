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
    message: str


def generate_link(code: str) -> str:
    if '' == code:
        return ''
    return f'{PROTOCOL}://{ROOT_URL}/download/{code}'


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = event.get('code', '')
    return {
        'success': '' != code,
        'code': code,
        'link': generate_link(code),
        'message': '',
    }
