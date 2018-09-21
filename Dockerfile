FROM golang:1.11.0 AS builder

RUN go get github.com/kr/godep

RUN mkdir -p /go/src/github.com/sh-miyoshi/kube-nodeport-checker
WORKDIR /go/src/github.com/sh-miyoshi/kube-nodeport-checker

COPY Godeps Godeps
COPY main.go main.go

RUN godep get

RUN go build -o kube-nodeport-checker main.go



FROM ubuntu:18.04
COPY --from=builder /go/src/github.com/sh-miyoshi/kube-nodeport-checker/kube-nodeport-checker /usr/local/bin

CMD ["kube-nodeport-checker", "--help"]
