# GinAdapter Lambda

## Overview
GinAdapter Lambda implemented AWS Lambda Go API Proxy integration for Gin framework. It allows you to run your Gin applications on AWS Lambda and API Gateway with minimal configuration.

## Features
- Easy integration with AWS Lambda and API Gateway
- Support for Gin middleware
- Automatic request and response handling

## Usage
1. Clone the repository:
   ```bash
    git clone git@github.com:InspectorGadget/ginadapter-lambda.git
    cd ginadapter-lambda
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Deploy the Lambda function:
    ```bash
    make create-function
    ```
