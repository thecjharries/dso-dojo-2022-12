from check_code_form.lambda_function import check_code_form
from check_code_used.lambda_function import check_code_used
from generate_link.lambda_function import generate_link


def test_check_code_form():
    assert 'abcde-12345-fghij' == check_code_form('abcde-12345-fghij')
    assert 'abcde-12345-fghij' == check_code_form('ABCDE-12345-FGHIJ')
    assert '' == check_code_form('test')


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
