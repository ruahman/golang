### debug
FROM golang:1.16-buster as debug

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /app
COPY . /app

RUN go build -o app main.go

RUN chmod +x dlv.sh 
ENTRYPOINT [ "dlv.sh"]

### dev
FROM golang:1.16-buster as dev

WORKDIR /app

### production
FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app