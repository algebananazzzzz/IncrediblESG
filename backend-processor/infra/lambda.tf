locals {
  function_name       = "${var.env}-app-func-${var.project_code}"
  execution_role_name = "${var.env}-app-role-${var.project_code}"
}


data "aws_ssm_parameter" "address" {
  name = "/shd/app/globalvpc/globalrediscache/ADDRESS"
}

module "function_execution_role" {
  source = "../../terraform-modules/iam_role"
  name   = local.execution_role_name
  policy_attachments = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
  ]
  custom_policy = {
    name = "${var.env}-app-policy-${var.project_code}"
    statements = {
      allowCreateNetworkInterface = {
        effect = "Allow"
        actions = [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface",
          "ec2:AssignPrivateIpAddresses",
          "ec2:UnassignPrivateIpAddresses"
        ]
        resources = ["*"]
      }
    }
  }
}

module "lambda_function" {
  source        = "../../terraform-modules/lambda_function"
  function_name = local.function_name
  runtime       = "provided.al2"
  handler       = "bootstrap"
  environment_variables = {
    REDIS_ADDR = "${data.aws_ssm_parameter.address.value}:6379"
    REDIS_KEY  = var.project_code
  }
  execution_role_arn = module.function_execution_role.role.arn
  deployment_package = {
    filename         = data.archive_file.deployment_package.output_path
    source_code_hash = data.archive_file.deployment_package.output_base64sha256
  }

  vpc_config = {
    subnet_ids         = data.aws_subnets.private.ids
    security_group_ids = [data.aws_security_group.allow_nat.id]
  }
}

data "archive_file" "deployment_package" {
  type        = "zip"
  source_dir  = "${path.module}/${var.deployment_package_location}"
  output_path = "${path.module}/../deploy/${var.project_code}.zip"
}
