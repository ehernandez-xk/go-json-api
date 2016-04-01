FROM golang

MAINTAINER Eddy Hernández <eddy.82.h@gmail.com>

RUN mkdir /usr/src/myapp
RUN go get github.com/gorilla/mux

WORKDIR /usr/src/myapp

COPY . .

EXPOSE 8080

# To build using this dockerfile
# docker build -t eddygt/json-api .
# I already pushed to docker hub
# docker push eddygt/json-api

# docker run -d -p 8080:8080 eddygt/json-api go run *.go
