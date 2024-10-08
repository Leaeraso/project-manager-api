# syntax=docker/dockerfile:1

# Use the appropriate Go version
FROM golang:1.22.4 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
WORKDIR /

COPY --from=build-stage /api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
