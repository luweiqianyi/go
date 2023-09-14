FROM golang:1.20 as builder
MAINTAINER runningriven@gmail.com

WORKDIR /usr/src/metrics/
COPY . .
ENV GOPROXY=https://goproxy.cn,direct

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /metrics/bin/service2/ ./cmd/service2

FROM alpine
WORKDIR /usr/src/metrics/
COPY --from=builder /metrics/bin/service2 ./
ENTRYPOINT ["./service2"]