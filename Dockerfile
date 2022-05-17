FROM golang:1.18-alpine

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /container-status

EXPOSE 80

CMD [ "/container-status" ]
