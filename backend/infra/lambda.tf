locals {
  execution_role_name = "${var.env}-mgmt-iamrole-${var.project_code}"
  function_name       = "${var.env}-app-func-${var.project_code}"
}


module "function_execution_role" {
  source = "../../terraform-modules/iam_role"
  name   = local.execution_role_name
  policy_attachments = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

module "lambda_function" {
  source             = "../../terraform-modules/lambda_function"
  function_name      = local.function_name
  execution_role_arn = module.function_execution_role.role.arn
  deployment_package = {
    filename         = data.archive_file.deployment_package.output_path
    source_code_hash = data.archive_file.deployment_package.output_base64sha256
  }
  runtime = "nodejs18.x"
  handler = "index.handler"
}

data "archive_file" "deployment_package" {
  type        = "zip"
  source_dir  = "${path.module}/${var.deployment_package_location}"
  output_path = "${path.module}/../deploy/${var.project_code}.zip"
}
