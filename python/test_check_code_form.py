from check_code_form import check_code_form


def test_check_code_form():
    assert 'abcde-12345-fghij' == check_code_form('abcde-12345-fghij')
    assert 'abcde-12345-fghij' == check_code_form('ABCDE-12345-FGHIJ')
    assert '' == check_code_form('test')
