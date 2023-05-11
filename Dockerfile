FROM golang:1.19.5 AS builder
# smoke test to verify if golang is available
RUN go version