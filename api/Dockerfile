FROM golang:1.24 AS base
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]

FROM base AS debug
RUN go mod download
RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air"]

FROM base AS build
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -o main .

FROM alpine:3.21 AS release
WORKDIR /root/
COPY --from=build /app/main .
CMD ["./main"]