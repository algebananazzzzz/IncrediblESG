variable "aws_region" {
  description = "The AWS region where resources will be deployed."
  default     = "ap-southeast-1"
}

variable "project_code" {
  description = "The code name of the project used for naming convention."
}

variable "env" {
  description = "The target environment to which the resources will be deployed."
}

variable "deployment_package_location" {
  description = "The location of the deployment package containing the Lambda code."
}

variable "custom_domain_name" {
  description = "The custom domain name for the API. Specify this variable to create a custom domain mapping for the API gateway."
  default     = null
}

variable "regional_certificate_domain" {
  description = "The domain name associated with the regional certificate. Specify this only if `custom_domain_name` is specified."
  default     = null
}

variable "route53_zone_name" {
  description = "The name of the Route 53 hosted zone to create DNS record for custom domain name. Specify this only if `custom_domain_name` is specified."
  default     = null
}
