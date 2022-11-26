remote_state {
  backend = "s3"
  generate = {
    path      = "generated.backend.tf"
    if_exists = "overwrite"
  }
  config = {
    bucket         = "wotw-dojo-state"
    key            = "2022/12/terraform.tfstate"
    region         = "us-west-2"
    encrypt        = true
  }
}
