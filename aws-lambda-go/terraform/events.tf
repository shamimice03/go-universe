# EventBridge Rule for CodePipeline State Changes
resource "aws_cloudwatch_event_rule" "codepipeline_state_change" {
  name        = "capture-codepipeline-state-changes"
  description = "Capture CodePipeline execution state changes (SUCCEEDED or FAILED)"

  event_pattern = jsonencode({
    source      = ["aws.codepipeline"],
    detail-type = ["CodePipeline Pipeline Execution State Change"],
    detail = {
      state = ["SUCCEEDED", "FAILED"]
    }
  })
}

# Target to connect the EventBridge rule to the Lambda function
resource "aws_cloudwatch_event_target" "invoke_lambda" {
  rule      = aws_cloudwatch_event_rule.codepipeline_state_change.name
  target_id = "SendToLambda"
  arn       = module.go_lambda_function.lambda_function_arn
}
