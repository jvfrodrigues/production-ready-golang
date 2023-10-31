FROM golang:1.21

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install build-essential -y && \
    go install github.com/spf13/cobra-cli@v1.3.0 &&

CMD ["tail", "-f", "/dev/null"]