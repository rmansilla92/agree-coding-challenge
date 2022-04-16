FROM golang:latest

ENV APPLICATION_PACKAGE="./cmd/api"

RUN apt-get update; apt-get install sqlite3
