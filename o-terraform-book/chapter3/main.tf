# resource "aws_s3_bucket" "terraform_state" {
#   bucket = "terraform-up-and-running-state-20240304"

#   lifecycle {
#     prevent_destroy = true
#   }
# }

# resource "aws_s3_bucket_versioning" "enable" {
#   bucket = aws_s3_bucket.terraform_state.id
#   versioning_configuration {
#     status = "Enabled"
#   }
# }

# resource "aws_s3_bucket_server_side_encryption_configuration" "default" {
#   bucket = aws_s3_bucket.terraform_state.id
#   rule {
#     apply_server_side_encryption_by_default {
#       sse_algorithm = "AES256"
#     }
#   }
# }

# resource "aws_s3_bucket_public_access_block" "public_access" {
#   bucket                  = aws_s3_bucket.terraform_state.id
#   block_public_acls       = true
#   block_public_policy     = true
#   ignore_public_acls      = true
#   restrict_public_buckets = true
# }

provider "aws" {
  region = "us-west-2"
}

terraform {
  backend "s3" {
    bucket  = "terraform-up-and-running-state-20240304"
    key     = "workspace-example/terraform.tfstate"
    region  = "us-west-2"
    encrypt = true
  }
}

resource "aws_instance" "example" {
  ami           = "ami-052c9ea013e6e3567"
  instance_type = terraform.workspace == "default" ? "t2.micro" : "t2.medium"
}

