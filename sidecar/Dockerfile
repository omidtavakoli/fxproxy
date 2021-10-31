FROM golang:1.15 as builder

WORKDIR /go/src/fxproxy
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM builder as server_builder

WORKDIR /go/src/fxproxy
COPY . .
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/*.go

RUN BUILD_TIME=$(date +%Y/%m/%d-%H:%M:%S) \
 && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.BuildTime=$BUILD_TIME" -o server cmd/server/*.go


FROM debian:stretch-slim
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
WORKDIR /usr/local/

ENV APP_MODE="production"
ENV GIN_MODE=release
RUN echo "app mode is ${APP_MODE}"

COPY --from=server_builder /go/src/fxproxy/server .
RUN chmod +x server
COPY --from=server_builder /go/src/fxproxy/configs ./configs

CMD ["./server"]

