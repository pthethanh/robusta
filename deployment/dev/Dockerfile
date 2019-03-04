# Dockerfile for lnquy/realize (https://cloud.docker.com/repository/docker/lnquy/realize)

FROM golang:1.10.8-alpine3.9

ENV REALIZE_VERSION=2.0.2

RUN apk --no-cache add --virtual build-dependencies git \
    && git config --global http.https://gopkg.in.followRedirects true

# Compile realize
RUN go get -u -v github.com/oxequa/realize \
    && go get -u -v github.com/go-siris/siris/core/errors
RUN cd /go/src/github.com/oxequa/realize \
    && git checkout v${REALIZE_VERSION} \
    && go get github.com/oxequa/realize

# Clean up
RUN apk del --purge -r build-dependencies \
    && rm -rf /go/src/*

ENTRYPOINT ["realize"]

#FROM alpine:3.9
#RUN apk --no-cache add gcc musl-dev
#COPY --from=builder /go/bin/realize /home/realize
##WORKDIR /home/
#RUN chmod +x /home/realize
#ENV PATH="/home/:${PATH}"
#ENTRYPOINT ["realize"]
