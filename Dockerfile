FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o homie

FROM alpine:latest
ENV PORT=8080
EXPOSE 8080
WORKDIR /app
RUN addgroup -S homie && adduser -S homie -G homie
COPY --chown=homie:homie --from=builder /app/homie .
USER homie
CMD ["/app/homie"]
