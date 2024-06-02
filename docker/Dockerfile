#FROM golang:1.22 AS builder
#
## Set the working directory inside the builder container
#WORKDIR /app
#
## Copy the Go modules manifests
#COPY ./res/applerefurbished /app/

# Build the Go binary statically
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o applerefurbished src/main.go
#
## Use a distroless base image for the final stage
FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY ./res/applerefurbished /app/

# Copy the statically built Go binary from the builder
#COPY --from=builder /app/applerefurbished /app/applerefurbished

# Specify the entrypoint
ENTRYPOINT ["/app/applerefurbished"]
