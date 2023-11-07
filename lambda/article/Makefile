mocks:
	# Run mockgen to generate mock interfaces
	mockgen -source=./internal/controller/controller.go -destination=mocks/mock_controller.go -package=mocks
	mockgen -source=./pkg/utils/utils.go -destination=mocks/mock_utils.go -package=mocks
	# Print a message indicating the process is complete
	echo "Mock interfaces generated successfully."

createlambdafnction:
	aws lambda create-function \
    --function-name article \
    --runtime go1.x \
    --handler main \
    --role arn:aws:iam::996985152674:role/article \
    --zip-file fileb://main.zip

listiamrole:
	aws iam list-roles

createiamrole:
	aws iam create-role \
    --role-name article \           
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

updatelambda:
	aws lambda update-function-code \
    --function-name article \
    --zip-file fileb://main.zip