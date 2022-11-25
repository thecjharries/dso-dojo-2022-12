from aws_lambda_powertools import Logger
from aws_lambda_powertools.utilities.typing import LambdaContext
from typing import TypedDict
from boto3 import client as boto3_client
from json import dumps as json_dumps, load as json_load

logger = Logger()
client = boto3_client('lambda')


class InputEvent(TypedDict):
    code: str


class OutputEvent(TypedDict):
    success: bool
    code: str
    link: str
    message: str


def lambda_handler(event: InputEvent, context: LambdaContext) -> OutputEvent:
    logger.info(event)
    code = event['code']
    payload = {
        'code': code.lower(),
    }
    response = client.invoke(
        FunctionName='check_code_form',
        InvocationType='RequestResponse',
        Payload=json_dumps(payload),
    )
    response_payload = json_load(response['Payload'])
    if not response_payload['success']:
        return {
            'success': False,
            'code': code,
            'link': '',
            'message': 'Malformed code',
        }
    response = client.invoke(
        FunctionName='check_code_used',
        InvocationType='RequestResponse',
        Payload=json_dumps(payload),
    )
    response_payload = json_load(response['Payload'])
    if not response_payload['success']:
        return {
            'success': False,
            'code': code,
            'link': '',
            'message': 'Code already used',
        }
    response = client.invoke(
        FunctionName='generate_link',
        InvocationType='RequestResponse',
        Payload=json_dumps(payload),
    )
    response_payload = json_load(response['Payload'])
    return {
        'success': False,
        'code': code,
        'link': response_payload['link'],
        'message': 'Success',
    }
