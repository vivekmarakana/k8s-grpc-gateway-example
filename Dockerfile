# Builder Image
FROM golang:1.13.12-alpine3.12 as builder
RUN apk update && apk add make g++
WORKDIR /go/src/gateway
COPY . .
RUN make

# Runner Image
FROM alpine:3.12
COPY --from=builder /go/src/gateway/bin /go/bin/
# Gateway port
EXPOSE 8080
# gRPC port
EXPOSE 9090
CMD ["/go/bin/server"]
