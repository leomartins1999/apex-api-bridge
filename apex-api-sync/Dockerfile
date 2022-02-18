# https://docs.docker.com/language/golang/build-images/
FROM golang:1.17.7-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY models ./models
COPY deserializers ./deserializers
COPY *.go ./

RUN go build -o /apex-api-sync

CMD [ "/apex-api-sync" ]
