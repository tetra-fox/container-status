FROM golang:1.22-alpine AS build

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o ./container-info

FROM scratch

COPY --from=build /app/container-info /usr/bin/container-info

EXPOSE 8080
ENTRYPOINT ["/usr/bin/container-info"]