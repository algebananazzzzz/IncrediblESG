resource "aws_dynamodb_table" "metric_table" {
  name         = "${var.env}-db-dynamodbtable-${var.project_code}-metric"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "MetricId"

  attribute {
    name = "MetricId"
    type = "S"
  }
}

resource "aws_dynamodb_table" "user_table" {
  name         = "${var.env}-db-dynamodbtable-${var.project_code}-user"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"

  attribute {
    name = "UserId"
    type = "S"
  }
}
