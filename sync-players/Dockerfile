# https://docs.docker.com/language/golang/build-images/
# build image
FROM golang:1.17.7-alpine as build

WORKDIR /app/core

COPY ./apex-sync-commons /app/apex-sync-commons
COPY ./sync-players /app/core/

RUN go build -o /sync-players

# run image
FROM golang:1.17.7-alpine

COPY --from=build /sync-players /sync-players

CMD [ "/sync-players" ]
