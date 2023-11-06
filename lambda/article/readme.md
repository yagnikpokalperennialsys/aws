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
Go to the console.aws.amazon.com
---
- Create account
- Create a lambda function on the AWS
Update the handler from by default hello to main
---

Query parameter
---
```
curl --location 'https://6srzklwppa.execute-api.eu-central-1.amazonaws.com/getarticlequeryparam/articles?limit=2'
```
---
![Alt text](docs/image.png)
---
Path parameter
---
```
curl --location --request POST 'https://6srzklwppa.execute-api.eu-central-1.amazonaws.com/getarticlequeryparam/articles/1'
```
---
![!\[Alt text\](/docs/image-1.png)](docs/image-1.png)
---
API Gateway routes
![Alt text](docs/image-2.png)
--
Single lambda function used for both the routes
![Alt text](docs/image-3.png)
---
GIN framework with the lambda 
![Alt text](docs/image-4.png)
---
Referance
---
https://github.com/awslabs/aws-lambda-go-api-proxy/tree/master/sample

https://github.com/yagnikpokalperennialsys/blogsystem

---



