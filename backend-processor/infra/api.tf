locals {
  api_gateway_name = "${var.env}-web-apigw-${var.project_code}"
}

module "api" {
  source = "../../terraform-modules/apigw"
  name   = local.api_gateway_name
}

module "api_lambda_integration" {
  source                    = "../../terraform-modules/apigw_lambda_integration"
  function_name             = module.lambda_function.function.function_name
  function_invoke_arn       = module.lambda_function.function.invoke_arn
  api_gateway_id            = module.api.api.id
  api_gateway_execution_arn = module.api.api.execution_arn
}
