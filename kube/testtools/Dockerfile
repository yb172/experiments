FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/yb172/experiments/kube/testtools
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
COPY --from=builder /go/src/github.com/yb172/experiments/kube/testtools/app app
COPY --from=builder /go/src/github.com/yb172/experiments/kube/testtools/config.yaml config.yaml
ENTRYPOINT ["./app", "load", "-k=false"]
