FROM golang:latest 

RUN mkdir -p $GOPATH/src

ADD . $GOPATH/src/

WORKDIR  $GOPATH/src/backend

RUN go build -o app *.go

CMD ["./app"]
