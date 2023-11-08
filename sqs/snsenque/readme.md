To generate the the arm 64 binary
---
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
```

To make the zip file
--
```
zip main.zip ./main
```

---

IAM policy to send message
---
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "sqs:SendMessage",
            "Resource": "arn:aws:sqs:region:996985152674:article"
        }
    ]
}
```

Attatch IAM policy
---
```
aws iam create-role \
    --role-name sendmessagequeue \
    --assume-role-policy-document '{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "lambda.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }'

# Attach the policy for sending messages to SQS
aws iam attach-role-policy \
    --policy-arn arn:aws:iam::aws:policy/AmazonSQSFullAccess \
    --role-name sendmessagequeue

```

Create the lambda function
```
createlambdafunction:
	aws lambda create-function \
    --function-name sendmessagequeue \
    --runtime go1.x \
    --handler main \
    --role arn:aws:iam::996985152674:role/sendmessagequeue \
    --zip-file fileb://main.zip
```

Update the lambda function
``` 
aws lambda update-function-code \
    --function-name sendmessagequeue \
    --zip-file fileb://main.zip
```
Referance
---
https://github.com/devtopics-yt/aws-lambda-sqs

---
