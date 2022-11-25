from check_code_form.lambda_function import check_code_form, lambda_handler as check_code_form_lambda_handler
from check_code_used.lambda_function import check_code_used, lambda_handler as check_code_used_lambda_handler
from generate_link.lambda_function import generate_link, lambda_handler as generate_link_lambda_handler

from json import loads as json_loads
from pytest import fixture, mark

with open('events.json', 'r') as events_file:
    EVENTS = json_loads(events_file.read())


@fixture
def events():
    return EVENTS


@fixture
def context():
    return object()


def test_check_code_form():
    assert 'abcde-12345-fghij' == check_code_form('abcde-12345-fghij')
    assert 'abcde-12345-fghij' == check_code_form('ABCDE-12345-FGHIJ')
    assert '' == check_code_form('test')


@mark.parametrize('event_name', ['empty', 'valid_unused'])
def test_check_code_form_lambda_handler(events, event_name, context):
    event = events[event_name]['event']
    response = check_code_form_lambda_handler(event, context)
    assert response['success'] == events[event_name]['response']['success']
    assert response['code'] == events[event_name]['response']['code']


def test_check_code_used():
    assert 'abcde-12345-fghji' == check_code_used('abcde-12345-fghji')
    assert 'abcde-12345-fgijh' == check_code_used('abcde-12345-fgijh')
    assert '' == check_code_used('abcde-12345-fghij')


def test_generate_link():
    assert 'https://reallycoolwebsite.com/download/abcde-12345-fghij' == generate_link(
        'abcde-12345-fghij')
    assert 'https://reallycoolwebsite.com/download/abcde-12345-fghji' == generate_link(
        'abcde-12345-fghji')
    assert 'https://reallycoolwebsite.com/download/abcde-12345-fgijh' == generate_link(
        'abcde-12345-fgijh')
