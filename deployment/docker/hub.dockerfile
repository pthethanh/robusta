# build the server
FROM golang:alpine as gobuilder

RUN mkdir -p /build/src/github.com/pthethanh/robusta
ENV GOPATH=/build
WORKDIR /build/src/github.com/pthethanh/robusta
ADD vendor ./vendor
ADD internal ./internal
ADD main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o robusta .
# build the web
FROM node:alpine as webbuilder
RUN mkdir /build
ADD ./web /build
WORKDIR /build
RUN yarn install
RUN yarn build
# now we are
FROM alpine
RUN adduser -S -D -H -h /app robusta
COPY --from=gobuilder /build/src/github.com/pthethanh/robusta/robusta /app/
COPY --from=webbuilder /build/dist /app/web
WORKDIR /app
RUN chown -R robusta /app
RUN chmod +x robusta
EXPOSE 8080
USER robusta
CMD ["./robusta"]
