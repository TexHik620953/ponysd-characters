FROM golang:1.25.5-alpine as builder

ARG GOOSE_VERSION=latest
RUN go install github.com/pressly/goose/v3/cmd/goose@${GOOSE_VERSION}

FROM alpine

COPY --from=builder /go/bin/goose /usr/local/bin

WORKDIR /data

ENTRYPOINT [ "/usr/local/bin/goose"]

CMD [ "status" ]
