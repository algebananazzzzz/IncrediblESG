## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=1.4.2 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~>5.31.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_api"></a> [api](#module\_api) | ./modules/apigw | n/a |
| <a name="module_api_domain_integration"></a> [api\_domain\_integration](#module\_api\_domain\_integration) | ./modules/apigw_domain_integration | n/a |
| <a name="module_api_lambda_integration"></a> [api\_lambda\_integration](#module\_api\_lambda\_integration) | ./modules/apigw_lambda_integration | n/a |
| <a name="module_function_execution_role"></a> [function\_execution\_role](#module\_function\_execution\_role) | ./modules/iam_role | n/a |
| <a name="module_lambda_function"></a> [lambda\_function](#module\_lambda\_function) | ./modules/lambda_function | n/a |

## Resources

| Name | Type |
|------|------|
| [archive_file.deployment_package](https://registry.terraform.io/providers/hashicorp/archive/latest/docs/data-sources/file) | data source |
| [aws_acm_certificate.issued](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/acm_certificate) | data source |
| [aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | The AWS region where resources will be deployed. | `string` | `"ap-southeast-1"` | no |
| <a name="input_custom_domain_name"></a> [custom\_domain\_name](#input\_custom\_domain\_name) | The custom domain name for the API. Specify this variable to create a custom domain mapping for the API gateway. | `any` | `null` | no |
| <a name="input_deployment_package_location"></a> [deployment\_package\_location](#input\_deployment\_package\_location) | The location of the deployment package containing the Lambda code. | `any` | n/a | yes |
| <a name="input_env"></a> [env](#input\_env) | The target environment to which the resources will be deployed. | `any` | n/a | yes |
| <a name="input_project_name"></a> [project\_name](#input\_project\_name) | The code name of the project used for naming convention. | `any` | n/a | yes |
| <a name="input_regional_certificate_domain"></a> [regional\_certificate\_domain](#input\_regional\_certificate\_domain) | The domain name associated with the regional certificate. Specify this only if `custom_domain_name` is specified. | `any` | `null` | no |
| <a name="input_route53_zone_name"></a> [route53\_zone\_name](#input\_route53\_zone\_name) | The name of the Route 53 hosted zone to create DNS record for custom domain name. Specify this only if `custom_domain_name` is specified. | `any` | `null` | no |

## Outputs

No outputs.
