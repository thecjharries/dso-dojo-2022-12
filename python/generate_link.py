PROTOCOL = 'https'
ROOT_URL = 'reallycoolwebsite.com'


def generate_link(code: str) -> str:
    return f'{PROTOCOL}://{ROOT_URL}/download/{code}'
