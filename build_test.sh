#!/bin/bash

# cd
echo "cd ./"
cd ./


# build
echo "GOOS=linux GOARCH=amd64 go build -o admin main.go"
GOOS=linux GOARCH=amd64 go build -o admin main.go

# upload
echo "scp -i ~/.ssh/green-dynamics admin  ec2-user@ec2-3-106-203-17.ap-southeast-2.compute.amazonaws.com:/home/ec2-user/go/admin"
scp -i ~/.ssh/green-dynamics admin  ec2-user@ec2-3-106-203-17.ap-southeast-2.compute.amazonaws.com:/home/ec2-user/go/admin