### development
FROM golang:1.18-buster as development

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

WORKDIR /go/src/work

### debug
FROM development as debug

EXPOSE 2345

COPY . /go/src/work

RUN go get github.com/go-delve/delve/cmd/dlv
RUN go install github.com/go-delve/delve/cmd/dlv

RUN go build -o app main.go


# ENTRYPOINT dlv debug --headless --log -l 0.0.0.0:2345 --api-version=2

# COPY ./dlv.sh /go/src/work/
# RUN chmod +x dlv.sh 
# ENTRYPOINT [ "./dlv.sh"]

### production
FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app