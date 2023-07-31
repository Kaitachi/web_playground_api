# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /godocker
EXPOSE 8080
CMD [ "/godocker" ]
