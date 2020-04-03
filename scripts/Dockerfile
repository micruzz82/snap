FROM golang:latest

ENV GOPATH=$GOPATH:/go
COPY ./ /go/src/github.com/micruzz82/snap
RUN apt-get update \
    && rm -rf /var/lib/apt/lists/* \
    && git clone https://github.com/micruzz82/gomit.git /go/src/github.com/micruzz82/gomit \
    && /go/src/github.com/micruzz82/snap/scripts/deps.sh \
    && make -C /go/src/github.com/micruzz82/snap
WORKDIR /go/src/github.com/micruzz82/snap

