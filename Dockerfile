FROM golang:1.21.6-alpine3.19

# RUN apt-get update 
# RUN apt-get install software-properties-common -y

RUN apk update

RUN apk add bash
RUN apk add openjdk8 java-cacerts

RUN update-ca-certificates -f

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /user_app

COPY ./Liquibase /app
COPY run.bash /app
RUN sed -i 's/\r$//'  /app/run.bash

EXPOSE 8080