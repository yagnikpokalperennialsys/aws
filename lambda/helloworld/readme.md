To generate the the arm 64 binary
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
```

To make the zip file

```
zip main.zip ./main
```
Go to the console.aws.amazon.com
- Create account
- Create a lambda function on the AWS
Update the handler from by default hell to main
![Alt text](image-2.png)

Run the lambda into the AWS
![Alt text](image-1.png)