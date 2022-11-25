# DSO Dojo 2022-12 Lambdas

## Usage

1. Set up the policy and role as defined in [calling Lambdas from Lambdas][1]
2. Run `make` in this directory to build all the Lambdas
3. For each directory here, create a Lambda function with the same name as the directory
4. Open up `runner` and create a test case with the following data:

    ```json
    {
        "code": "abcde-12345-fgijh"
    }
    ```

5. Run the test case and see the output

[1]: <https://www.sqlshack.com/calling-an-aws-lambda-function-from-another-lambda-function/> "Calling Lambdas from Lambdas"
