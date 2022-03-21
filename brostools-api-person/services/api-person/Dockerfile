FROM golang:1.15.0-alpine

RUN apk add --no-cache \
        alpine-sdk \
        git


WORKDIR /go/src/brostools-api-person/
COPY ./services/api-person/ /go/src/brostools-api-person/

RUN go build

EXPOSE 8080

RUN go get github.com/pilu/fresh
CMD ["fresh"]