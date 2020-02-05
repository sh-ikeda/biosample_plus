FROM golang:1.13.7-alpine3.11 AS build-env

RUN apk --no-cache add git

RUN go get github.com/garyburd/redigo/redis && go get github.com/labstack/echo

WORKDIR /go/build

COPY . .
WORKDIR ./cmd/bspsrv/
RUN go build


FROM alpine:3.11

COPY --from=build-env /go/build/cmd/bspsrv/bspsrv /usr/local/bin
EXPOSE 8080
CMD ["bspsrv"]
