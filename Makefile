.PHONY: deps infra-init infra-apply infra-destroy

deps:
	which terraform

infra-init: deps
	terraform init

infra-apply: deps
	terraform apply \
		-var "aws_profile=$AWS_PROFILE" \
		-var "aws_region=$AWS_REGION" \
		-var "aws_event_stream_name=$AWS_EVENT_STREAM_NAME" \
		-var "serverless_user=$SERVERLESS_USER"

infra-destroy: deps
	terraform destroy \
		-var "aws_profile=$AWS_PROFILE" \
		-var "aws_region=$AWS_REGION" \
		-var "aws_event_stream_name=$AWS_EVENT_STREAM_NAME" \
		-var "serverless_user=$SERVERLESS_USER"
