# Code Flow

Codes are of the form `ABCDE-12345-FGHIJ`. Codes must be validated by checking form and lowercasing then ensuring they have not been used before. If these criteria are met, a link can be generated.

1. Check the code form [via `check_code_form()`](./check_code_form.py)
2. Check if the code has been used [via `check_code_used()`](./check_code_used.py)
3. Generate a link [via `generate_link()`](./generate_link.py)
