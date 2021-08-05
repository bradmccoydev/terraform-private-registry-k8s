FROM golang:1.15-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/server
RUN go build -o /go/bin/server/server

FROM alpine:3.10
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR app
COPY --from=GO_BUILD /go/bin/server/ ./
EXPOSE 8080
CMD ./server