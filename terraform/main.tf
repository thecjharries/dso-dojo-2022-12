terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

data "aws_caller_identity" "current" {}

output "caller" {
  value = data.aws_caller_identity.current
}

resource "null_resource" "zip" {
  provisioner "local-exec" {
    command = "cd ${path.module}/../lambdas && make clean && make all"
  }
}
