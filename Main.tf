variable "aws_profile" {
    type = string
}

variable "aws_region" {
    type = string
}

variable "aws_event_stream_name" {
    type = string
}

variable "serverless_user" {
    type = string
}

provider "aws" {
    profile    = var.aws_profile
    region     = var.aws_region
}

resource "aws_kinesis_stream" "event_stream" {
  name             = var.aws_event_stream_name
  shard_count      = 1
  retention_period = 48

  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]

  tags = {
    project = format("%s%s", "sls-demo-", var.serverless_user)
    owner = var.serverless_user
  }
}
