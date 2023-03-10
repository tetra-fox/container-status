FROM golang:1.20-alpine

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /container-info

EXPOSE 8080

CMD [ "/container-info" ]
