FROM golang:1.10.4-alpine3.8

RUN go get github.com/kr/godep

RUN mkdir -p /go/src/github.com/sh-miyoshi/kube-nodeport-checker
WORKDIR /go/src/github.com/sh-miyoshi/kube-nodeport-checker

COPY Godeps Godeps
COPY main.go main.go

RUN godep get

CMD ["go", "run", "main.go", "--help"]
