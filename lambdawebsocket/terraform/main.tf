provider "aws" {
  profile = "jds"
  region = "us-west-1"
}


resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambdawebsocket_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "lambdawebsocket_lambda" {
  filename      = "../lambda.zip"
  function_name = "lambdawebsocketlambda"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "lambdawebsocket"
  source_code_hash = data.archive_file.lambda_zip.output_base64sha256
  runtime = "go1.x"
}

resource "aws_lambda_alias" "lambdawebsocket_lambda_live" {
  name             = "live"
  description      = "set a live alias"
  function_name    = aws_lambda_function.lambdawebsocket_lambda.arn
  function_version = aws_lambda_function.lambdawebsocket_lambda.version
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../lambdawebsocket"
  output_path = "../lambda.zip"
}

resource "aws_lambda_permission" "lambdawebsocket_lambda" {
  statement_id  = "AllowExecutionFromApplicationLoadBalancer"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambdawebsocket_lambda.arn
  principal     = "elasticloadbalancing.amazonaws.com"
  source_arn = aws_lb_target_group.lambdawebsocket_lambda.arn
}

resource "aws_lambda_permission" "lambdawebsocket_lambda_live" {
  statement_id  = "AllowExecutionFromApplicationLoadBalancer"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_alias.lambdawebsocket_lambda_live.arn
  principal     = "elasticloadbalancing.amazonaws.com"
  source_arn = aws_lb_target_group.lambdawebsocket_lambda.arn
}

# IAM
resource "aws_iam_role_policy_attachment" "cloudwatch-attach" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}


# ALB
resource "aws_lb_target_group" "lambdawebsocket_lambda" {
  name        = "lambdawebsocketlambda"
  target_type = "lambda"
}

resource "aws_lb_target_group_attachment" "lambdawebsocket_lambda" {
  target_group_arn  = aws_lb_target_group.lambdawebsocket_lambda.arn
  target_id         = aws_lambda_alias.lambdawebsocket_lambda_live.arn
  depends_on        = [aws_lambda_permission.lambdawebsocket_lambda_live]
}

resource "aws_lb_listener_rule" "poopjournal_server_lambda" {
  listener_arn = data.terraform_remote_state.stinkyfingers.outputs.stinkyfingers_https_listener
  priority = 98
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lambdawebsocket_lambda.arn
  }
  condition {
    path_pattern {
      values = ["/lambdawebsocket/*"]
    }
  }
  depends_on = [aws_lb_target_group.lambdawebsocket_lambda]
}

# backend
# data "terraform_remote_state" "poopjournal" {
#   backend = "s3"
#   config = {
#     bucket  = "remotebackend"
#     key     = "poopjournal/terraform.tfstate"
#     region  = "us-west-1"
#     profile = "jds"
#   }
# }
#
data "terraform_remote_state" "stinkyfingers" {
  backend = "s3"
  config = {
    bucket  = "remotebackend"
    key     = "stinkyfingers/terraform.tfstate"
    region  = "us-west-1"
    profile = "jds"
  }
}
