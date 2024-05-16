FROM golang:1.22 as builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 \
    go build -o tracy cmd/tracy/*.go

FROM alpine as runner
WORKDIR /app
COPY --from=builder /app/tracy ./