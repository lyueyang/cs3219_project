# syntax=docker/dockerfile:1

FROM golang:1.17.0-alpine

WORKDIR /app

# copy the files we need to get started
COPY ./ ./

# get dependencies we need
RUN go mod download

# build our project
RUN go build -o /cs3219_project

# run this on startup
CMD [ "/cs3219_project" ]