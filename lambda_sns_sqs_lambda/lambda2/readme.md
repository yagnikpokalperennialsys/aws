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
Referance
---
https://github.com/devtopics-yt/lambda-publish-sns-sqs

https://docs.aws.amazon.com/lambda/latest/dg/with-sqs-create-package.html

---



