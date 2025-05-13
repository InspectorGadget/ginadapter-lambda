.PHONY: build clean test deploy

# Binary name
BINARY_NAME=ginadapter-lambda
LAMBDA_HANDLER=bootstrap

# AWS Lambda settings
LAMBDA_FUNCTION_NAME=ginadapter-lambda-function
AWS_REGION=ap-southeast-1

# IAM Role
IAM_ROLE_ARN=arn:aws:iam::ACCOUNT_ID:role/lambda-role

# Go settings
GOFLAGS=-ldflags="-s -w"

build:
	env GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o ./bin/$(LAMBDA_HANDLER)

build-local:
	go build $(GOFLAGS) -o ./bin/$(BINARY_NAME) .

clean:
	rm -rf ./bin

zip: clean build
	zip ./bin/deployment.zip ./bin/$(LAMBDA_HANDLER)

deploy: zip
	aws lambda update-function-code \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--zip-file fileb://bin/deployment.zip \
		--region $(AWS_REGION)

create-function: zip
	aws lambda create-function \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--runtime provided.al2 \
		--handler $(LAMBDA_HANDLER) \
		--zip-file fileb://bin/deployment.zip \
		--role $(IAM_ROLE_ARN) \
		--region $(AWS_REGION)

run-local: build-local
	./bin/$(BINARY_NAME)