

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
Receive 2 messages of last 10 seconds
![img_1.png](img_1.png)