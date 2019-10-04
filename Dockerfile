FROM golang:1.13
RUN apt-get update && \
    apt-get -y install --no-install-recommends \
    redis-server
RUN go get github.com/garyburd/redigo/redis && go get github.com/labstack/echo
WORKDIR /usr/local/bin
COPY ./biosample_plus.go ./
RUN go build -o biosample_plus biosample_plus.go
EXPOSE 8080
CMD ["./biosample_plus"]
