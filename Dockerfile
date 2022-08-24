FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git
RUN apk add --update make

WORKDIR /tmp/go-databases

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

# Create a new fresh image, and optmize the size of the application
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-databases/out/go-databases /app/go-databases
# COPY --from=build_base /tmp/go-databases/.env /app/go-databases

EXPOSE 8080

CMD ["./app/go-databases"]