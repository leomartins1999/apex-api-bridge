# https://docs.docker.com/language/golang/build-images/
# build image
FROM golang:1.17.7-alpine as build

WORKDIR /app/core

COPY ./apex-sync-commons /app/apex-sync-commons
COPY ./sync-games /app/core/

RUN go build -o /sync-games

# run image
FROM golang:1.17.7-alpine

COPY --from=build /sync-games /sync-games

CMD [ "/sync-games" ]
