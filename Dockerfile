# Build stage #

FROM golang:1.21 AS builder

WORKDIR /opt

COPY . /opt

RUN go build -o ./runner ./cmd

# Deploy stage # 

FROM ubuntu:latest

WORKDIR /opt 

COPY --from=builder /opt/runner /opt/

ENTRYPOINT /opt/runner
