FROM golang:alpine as builder

RUN apk update && apk add ca-certificates && apk add make

COPY . $GOPATH/src/github.com/inquizarus/godnd/magic/

WORKDIR $GOPATH/src/github.com/inquizarus/godnd/magic
RUN CGO111MODULE=on make build-linux
RUN mv magic_unix /go/bin/magic

FROM busybox:latest
COPY --from=builder /go/bin/magic /go/bin/magic
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
EXPOSE 8080
WORKDIR /go/bin
CMD ./magic