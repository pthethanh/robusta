FROM alpine:3.8

WORKDIR /home/
COPY robusta.bin .
RUN chmod +x robusta.bin

COPY web/dist ./web/dist
COPY configs ./configs
COPY templates ./templates

EXPOSE 80
CMD ["./robusta.bin"]
