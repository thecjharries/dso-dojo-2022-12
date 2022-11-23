from generate_link import generate_link


def test_generate_link():
    assert 'https://nexient.com/download/abcde-12345-fghij' == generate_link(
        'abcde-12345-fghij')
    assert 'https://nexient.com/download/abcde-12345-fghji' == generate_link(
        'abcde-12345-fghji')
    assert 'https://nexient.com/download/abcde-12345-fgijh' == generate_link(
        'abcde-12345-fgijh')
