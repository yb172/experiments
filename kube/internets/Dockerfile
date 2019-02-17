FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/yb172/experiments/kube/internets
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
COPY --from=builder /go/src/github.com/yb172/experiments/kube/internets/app app
COPY --from=builder /go/src/github.com/yb172/experiments/kube/internets/config.yaml config.yaml
EXPOSE 10000
ENTRYPOINT ["./app"]
