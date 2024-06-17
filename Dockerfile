FROM golang:1.22-alpine AS build

ENV GIN_MODE=release

COPY go.mod /
COPY go.sum /
COPY *.go /

RUN go mod download
RUN go build -o /container-info

FROM alpine:latest

COPY --from=build /container-info /usr/bin/container-info

EXPOSE 8080
ENTRYPOINT ["/usr/bin/container-info"]