module "go_lambda_function" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "7.0.0"

  publish       = true
  function_name = "event-printer"

  attach_cloudwatch_logs_policy     = true
  cloudwatch_logs_retention_in_days = 1

  create_lambda_function_url = true

  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["x86_64"] # x86_64 (GOARCH=amd64); arm64 (GOARCH=arm64)

  trigger_on_package_timestamp = false

  source_path = [
    {
      path = "${path.module}/../src"
      commands = [
        "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go",
        ":zip",
      ]
      patterns = [
        "!.*",
        "bootstrap",
      ]
    }
  ]

  allowed_triggers = {
    eventbridge = {
      principal  = "events.amazonaws.com"
      source_arn = aws_cloudwatch_event_rule.codepipeline_state_change.arn
    }
  }
}
