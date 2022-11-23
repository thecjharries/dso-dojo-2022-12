from check_code_used import check_code_used
from itertools import permutations


def test_check_code_used():
    assert 'abcde-12345-fghji' == check_code_used('abcde-12345-fghji')
    assert 'abcde-12345-fgijh' == check_code_used('abcde-12345-fgijh')
    assert '' == check_code_used('abcde-12345-fghij')
