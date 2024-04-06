# locals {
#   firehose_stream_name      = "${var.env}-app-firehosestream-${var.project_code}"
#   firehose_exec_role_name   = "${var.env}-app-role-${var.project_code}-firehose"
#   firehose_exec_policy_name = "${var.env}-app-policy-${var.project_code}-firehose"
#   data_bucket_name          = "${var.env}-app-s3-${lower(var.project_code)}-databucket"
# }


# module "firehose_execution_role" {
#   source = "../../terraform-modules/iam_role"
#   name   = local.firehose_exec_role_name
#   assume_role_allowed_principals = [{
#     type        = "Service"
#     identifiers = ["firehose.amazonaws.com"]
#   }]
#   custom_policy = {
#     name = local.firehose_exec_policy_name
#     statements = {
#       allowS3Put = {
#         effect = "Allow"
#         actions = [
#           "logs:PutLogEvents",
#           "s3:AbortMultipartUpload",
#           "s3:GetBucketLocation",
#           "s3:GetObject",
#           "s3:ListBucket",
#           "s3:ListBucketMultipartUploads",
#           "s3:PutObject"
#         ]
#         resources = ["*"]
#       }
#     }
#   }
# }


# resource "aws_kinesis_firehose_delivery_stream" "extended_s3_stream" {
#   name        = local.firehose_stream_name
#   destination = "extended_s3"

#   extended_s3_configuration {
#     role_arn           = module.firehose_execution_role.role.arn
#     bucket_arn         = aws_s3_bucket.data_bucket.arn
#     buffering_interval = 900
#     buffering_size     = 128

#     dynamic_partitioning_configuration {
#       enabled = "true"
#     }
#     prefix              = "data/metric_id=!{partitionKeyFromQuery:metric_id}/"
#     error_output_prefix = "errors/!{firehose:error-output-type}/"

#     processing_configuration {
#       enabled = "true"

#       processors {
#         type = "MetadataExtraction"
#         parameters {
#           parameter_name  = "JsonParsingEngine"
#           parameter_value = "JQ-1.6"
#         }
#         parameters {
#           parameter_name  = "MetadataExtractionQuery"
#           parameter_value = "{metric_id:.metric_id}"
#         }
#       }
#     }
#   }
# }

# resource "aws_s3_bucket" "data_bucket" {
#   bucket = local.data_bucket_name
# }
