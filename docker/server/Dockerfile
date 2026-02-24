FROM golang:1.25.5-alpine AS base

WORKDIR /app
RUN apk add --no-cache git openssh

COPY . .
RUN go mod download
RUN go generate ./...
RUN go build -o ./bin/app ./cmd/server/


FROM alpine

COPY --from=base /app/bin/app /app

CMD [ "/app" ]