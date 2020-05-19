FROM golang:alpine as builder
WORKDIR /build
ADD . .
ARG opts
RUN env CGO_ENABLED=0 ${opts} go build -ldflags="-w -s" -o cfn-format ./cmd/cfn-format/*

FROM scratch
COPY --from=builder /build/cfn-format /go/bin/cfn-format
ENTRYPOINT ["/go/bin/cfn-format"]
