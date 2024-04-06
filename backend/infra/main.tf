provider "aws" {
  region = var.aws_region
}

terraform {
  required_version = ">=1.4.2"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>5.31.0"
    }
  }

  backend "s3" {
    bucket               = "shd-all-bucket-tfstate-algebananazzzzz"
    key                  = "incrediblESG-api.tfstate"
    workspace_key_prefix = "tf-state"
    region               = "ap-southeast-1"
  }
}
