FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/yb172/experiments/kube/gateway
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
COPY --from=builder /go/src/github.com/yb172/experiments/kube/gateway/app app
COPY --from=builder /go/src/github.com/yb172/experiments/kube/gateway/config.yaml config.yaml
EXPOSE 8000
ENTRYPOINT ["./app"]
