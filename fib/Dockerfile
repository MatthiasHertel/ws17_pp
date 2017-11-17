FROM golang

WORKDIR /go/src/fib
COPY . .

RUN go build . && mv ./fib /go/bin/

CMD ["fib"]
