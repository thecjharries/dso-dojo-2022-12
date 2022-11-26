terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = "dso-dojo-2022-12-lambda-basic"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_role_policy_attachment" "basic_execution" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

locals {
  lambdas = toset([
    "check_code_form",
    "check_code_used",
    "generate_link",
  ])
}

resource "null_resource" "lambdas" {
  for_each = setunion(local.lambdas, ["runner"])

  triggers = {
    code = file("${path.module}/../${each.key}/lambda_function.py")
  }

  provisioner "local-exec" {
    command = <<EOT
      cd ${path.module}/..
      rm -rf ${each.key}.zip
      cd ${each.key}
      rm -rf requirements.txt package
      pipenv requirements > requirements.txt
      pipenv install --target package --requirement requirements.txt
      cd package
      zip -r ../${each.key}.zip .
      cd ..
      zip -g ${each.key}.zip lambda_function.py
    EOT
  }
}

resource "aws_lambda_function" "basic_lambdas" {
  for_each = local.lambdas

  filename      = "${path.module}/../${each.key}.zip"
  function_name = each.key
  role          = aws_iam_role.lambda.arn
  runtime       = "python3.9"
  handler       = "lambda_function.lambda_handler"
}

data "aws_iam_policy_document" "lambda_invoke" {
  statement {
    actions = [
      "lambda:InvokeFunction",
      "lambda:InvokeAsync",
    ]

    resources = [
      for lambda in aws_lambda_function.basic_lambdas : lambda.arn
    ]
  }
}

resource "aws_iam_policy" "lambda_invoke" {
  name   = "dso-dojo-2022-12-lambda-basic-invoke"
  policy = data.aws_iam_policy_document.lambda_invoke.json
}

resource "aws_iam_role" "lambda_invoke" {
  name               = "dso-dojo-2022-12-lambda-basic-invoke"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_role_policy_attachment" "lambda_invoke" {
  role       = aws_iam_role.lambda_invoke.name
  policy_arn = aws_iam_policy.lambda_invoke.arn
}

resource "aws_iam_role_policy_attachment" "invoke_execution" {
  role       = aws_iam_role.lambda_invoke.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "runner" {
  filename      = "${path.module}/../runner.zip"
  function_name = "runner"
  role          = aws_iam_role.lambda_invoke.arn
  runtime       = "python3.9"
  handler       = "lambda_function.lambda_handler"
}

resource "aws_lambda_function_url" "runner" {
  function_name      = aws_lambda_function.runner.function_name
  authorization_type = "NONE"
}

output "runner_url" {
  value = aws_lambda_function_url.runner.function_url
}
