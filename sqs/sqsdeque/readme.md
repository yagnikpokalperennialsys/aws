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
Create the role
``` 
aws iam create-role \
    --role-name dequeuerole \
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
```
Create the iam policy
```
aws iam create-policy \
    --policy-name consumequeue-policy \
    --policy-document '{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "sqs:ReceiveMessage",
                    "sqs:DeleteMessage"
                ],
                "Resource": "arn:aws:sqs:region:996985152674:queue-name"
            }
        ]
    }'

```
Attatch the iam policy
```
aws iam attach-role-policy \
    --role-name consumequeue \
    --policy-arn arn:aws:iam::996985152674:policy/consumequeue-policy
```
Create the lambda function
```
aws lambda create-function \
    --function-name receivemessagefromqueue \
    --runtime go1.x \
    --handler main \
    --role arn:aws:iam::996985152674:role/consumequeue
    --zip-file fileb://main.zip

```
---
Referance
---
https://github.com/devtopics-yt/lambda-publish-sns-sqs

https://docs.aws.amazon.com/lambda/latest/dg/with-sqs-create-package.html

---



