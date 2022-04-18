FROM golang:1.15

ENV APPLICATION_PACKAGE="./cmd/api"
ENV PORT 8080 

RUN apt-get update; apt-get install sqlite3
