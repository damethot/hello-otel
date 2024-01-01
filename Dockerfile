FROM golang AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/otel otel

FROM scratch
COPY --from=builder /out/otel /app/otel

EXPOSE 3000
CMD ["/app/otel"]