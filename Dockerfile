FROM --platform=linux/amd64 docker.io/golang:1.21-alpine3.18 AS builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on

ADD . /app
WORKDIR /app

# Stage 1: Build the binary
RUN GOOS=linux GOARCH=amd64 go build -o supply-server .

# Stage 2: Create the runtime image
FROM --platform=linux/amd64 docker.io/alpine:3.18.4

RUN apk update && apk add --no-cache bash curl jq
COPY --from=builder /app/supply-server .

EXPOSE 8080
CMD ["./supply-server"]
