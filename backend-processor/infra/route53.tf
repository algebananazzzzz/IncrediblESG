locals {
  use_custom_domain = var.custom_domain_name != null
}

data "aws_acm_certificate" "issued" {
  count       = local.use_custom_domain ? 1 : 0
  domain      = var.regional_certificate_domain
  statuses    = ["ISSUED"]
  most_recent = true
}

data "aws_route53_zone" "zone" {
  count        = local.use_custom_domain ? 1 : 0
  name         = var.route53_zone_name
  private_zone = false
}

module "api_domain_integration" {
  count                    = local.use_custom_domain ? 1 : 0
  source                   = "../../terraform-modules/apigw_domain_integration"
  domain_name              = var.custom_domain_name
  regional_certificate_arn = data.aws_acm_certificate.issued[0].arn
  zone_id                  = data.aws_route53_zone.zone[0].id
  apigw_api_id             = module.api.api.id
  apigw_stage_id           = module.api.stage.id
}
