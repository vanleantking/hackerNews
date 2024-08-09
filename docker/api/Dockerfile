FROM golang:1.21.5 AS build-stage
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates
WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Deploy the application binary into a lean image
FROM golang:1.21.5 AS build-release-stage
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates 
WORKDIR /app

COPY --from=build-stage /app/main /app/main
COPY --from=build-stage /app/.env /app/.env
RUN ls -la
EXPOSE ${API_PORT}
ENTRYPOINT [ "./main"]