FROM golang

WORKDIR /go/src/fib
COPY . .

RUN go build fib.go && mv ./fib /go/bin/

CMD ["fib"]
