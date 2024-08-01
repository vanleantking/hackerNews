FROM golang:1.21.5 AS build-stage
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Deploy the application binary into a lean image
FROM golang:1.21.5 AS build-release-stage
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates 
WORKDIR /app

COPY --from=build-stage /app/main /app/main
EXPOSE 8080
ENTRYPOINT [ "./main"]