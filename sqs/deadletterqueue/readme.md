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
Create a queue
![img.png](img.png)

Create a deadletter queue
![img_1.png](img_1.png)

Update main queue with deadletter queue with 3 retry
![img_10.png](img_10.png)

Deploy lambda function

![img_11.png](img_11.png)

Subscribe the lamnda with SQS
![img_3.png](img_3.png)

Send message to SQS queue
![img_6.png](img_6.png)

- Got the 3 retry while checking the lambda log
![img_4.png](img_4.png)



Message comes but not consumes and delete by queue when lambda was triggered
![img_5.png](img_5.png)

Check the message on deadletter queue while polling
![img_9.png](img_9.png)

![img_8.png](img_8.png)