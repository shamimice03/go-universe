## 1. Build your Lambda binary for AWS
```
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go  
```

## 2. Create the deployment package
```
zip lambda-function.zip bootstrap
```

### 3. Update Function Code
```
aws lambda update-function-code \
  --function-name test-printer \
  --zip-file fileb://lambda-function.zip
```

### Script
```
#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Set variables
FUNCTION_NAME="test-printer"
REGION="ap-northeast-1"  

echo "Building Go binary for AWS Lambda..."
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

echo "Creating deployment package..."
zip lambda-function.zip bootstrap

echo "Updating Lambda function code..."
aws lambda update-function-code \
  --function-name $FUNCTION_NAME \
  --zip-file fileb://lambda-function.zip \
  --region $REGION

echo "Deployment completed successfully!"
echo "Function name: $FUNCTION_NAME"
echo "Region: $REGION"
```