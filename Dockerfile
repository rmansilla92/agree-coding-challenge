FROM golang:1.16-buster as builder

ENV APPLICATION_PACKAGE="./cmd/api"
ENV PORT 8080 

RUN apt-get update; apt-get install sqlite3
