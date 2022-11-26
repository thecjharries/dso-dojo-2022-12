# Putting it All Together

## Duplicated code?

Yes. This folder duplicates everything because I want to leave my iterations in the repo for people in the Dojo to see. That makes my life a little more difficult; it's a fun challenge.

## Goals

I'd like to create the following workflow:

* Test all the Python except for the runner using Pytest
* Deploy everything to LocalStack using Terragrunt (done)
* Test the runner using Terratest on LocalStack (done)
* Deploy everything to AWS using Terragrunt
