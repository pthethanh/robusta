FROM alpine:3.8

WORKDIR /home/
COPY robusta.bin .
RUN chmod +x robusta.bin

COPY web/dist ./web

EXPOSE 8080
CMD ["./robusta.bin"]
