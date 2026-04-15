FROM golang:1.23.1 AS builder

WORKDIR /app

COPY . .

RUN if [ ! -f go.mod ]; then go mod init github.com/santosbruno98/price-scraper; fi \
    && go mod tidy

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -a -o /output/price-scraper .

FROM alpine:3.19

RUN apk add --no-cache ca-certificates \
    && adduser -D -g '' appuser

WORKDIR /home/appuser/

COPY --from=builder /output/price-scraper .

RUN chown -R appuser:appuser /home/appuser \
    && chmod +x price-scraper

USER appuser

ENTRYPOINT ["./price-scraper"]