FROM golang:1.21-alpine as builder
WORKDIR /builder
COPY . .
RUN apk add --no-cache upx \
    && go mod download \
    && go build -ldflags "-s -w" -o main \
    && upx -9 main

FROM alpine:latest
ARG ENV
ENV ENV=${ENV}
WORKDIR /app
COPY --from=builder /builder/main .
COPY --from=builder /builder/migrations ./migrations
COPY --from=builder /builder/config ./config
CMD ["/app/main", "server"]