locals {
  allow_from_nat_instance_sg_name = "shd-app-sg-allownat"
}

data "aws_subnets" "private" {
  filter {
    name = "tag:Name"
    values = [
      "${var.env}-app-subnet-private-1a",
      "${var.env}-app-subnet-private-1b"
    ]
  }
}

data "aws_security_group" "allow_nat" {
  filter {
    name   = "tag:Name"
    values = [local.allow_from_nat_instance_sg_name]
  }
}
