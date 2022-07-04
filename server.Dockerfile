# syntax=docker/dockerfile:1
FROM golang:1.17

WORKDIR /GOOauth

COPY go.mod go.sum config.yml ./
RUN go mod download && go mod verify
COPY . .

RUN go build
CMD ["go" ,"run" ,"GOOauth"]
EXPOSE 8090
