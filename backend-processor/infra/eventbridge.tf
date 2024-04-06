locals {
  scheduler_execution_role_name        = "${var.env}-mgmt-role-${var.project_code}-scheduler"
  scheduler_execution_role_policy_name = "${var.env}-mgmt-policy-${var.project_code}-scheduler"
  scheduler_group_name                 = "${var.env}-app-schedulegrp-${var.project_code}"
  schedule_rate_5                      = "${var.env}-app-schedule-${var.project_code}-rate-5"
}

module "scheduler_execution_role" {
  source = "../../terraform-modules/iam_role"
  name   = local.scheduler_execution_role_name

  assume_role_allowed_principals = [{
    type        = "Service"
    identifiers = ["scheduler.amazonaws.com"]
  }]

  custom_policy = {
    name = local.scheduler_execution_role_policy_name
    statements = {
      allowLambdaInvoke = {
        effect = "Allow"
        actions = [
          "lambda:InvokeFunction",
        ]
        resources = [module.lambda_function.function.arn]
      }
    }
  }
}

resource "aws_scheduler_schedule_group" "this" {
  name = local.scheduler_group_name
}

resource "aws_scheduler_schedule" "rate_5" {
  name       = local.schedule_rate_5
  group_name = aws_scheduler_schedule_group.this.name

  flexible_time_window {
    mode = "OFF"
  }
  schedule_expression = "rate(5 minutes)"
  start_date          = "2024-04-06T11:00:00Z"

  target {
    arn      = module.lambda_function.function.arn
    role_arn = module.scheduler_execution_role.role.arn

    input = jsonencode({
      command = "reset_cache"
    })
  }

  lifecycle {
    ignore_changes = [start_date]
  }
}
