FROM golang:1.13.0-alpine3.10

ENV REALIZE_VERSION=2.0.2

ENV GO111MODULE=off

RUN apk --no-cache add --virtual build-dependencies git \
    && git config --global http.https://gopkg.in.followRedirects true

# Compile realize
RUN go get -u -v github.com/oxequa/realize \
    && go get -u -v github.com/go-siris/siris/core/errors
RUN cd /go/src/github.com/oxequa/realize \
    && git checkout v${REALIZE_VERSION} \
    && go get github.com/oxequa/realize

ENTRYPOINT ["realize"]
