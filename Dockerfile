FROM golang:1.10-alpine as builder

RUN apk add --no-cache make curl git gcc musl-dev linux-headers
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

ADD . /go/src/github.com/linkpoolio/alpha-vantage-cl-ea
RUN cd /go/src/github.com/linkpoolio/alpha-vantage-cl-ea && make build

# Copy Adaptor into a second stage container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/linkpoolio/alpha-vantage-cl-ea/alpha-vantage-cl-ea /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["alpha-vantage-cl-ea"]