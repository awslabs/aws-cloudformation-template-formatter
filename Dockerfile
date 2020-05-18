FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
ARG opts
RUN env ${opts} go build -ldflags="-w -s" -o cfn-format ./cmd/cfn-format/*

FROM scratch
COPY --from=builder /build/cfn-format /go/bin/cfn-format
ENTRYPOINT ["/go/bin/cfn-format"]
