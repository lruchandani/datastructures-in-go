#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/offersender
echo "Built Package"
docker build . -t lruchandani/offersender
echo "Build Image"
docker push lruchandani/offersender
echo "Pushed imaged"