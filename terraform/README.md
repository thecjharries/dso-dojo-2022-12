# Terraform for DSO Dojo 2022-12

## Usage

1. Change the state bucket in [`terragrunt.hcl`](./terragrunt.hcl) to your own bucket.
2. `cd ../lambdas && make clean && make && cd ../terraform`
3. `terragrunt init`
4. `terragrunt apply`
5. `./test.sh`
