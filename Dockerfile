# Build stage - Copy src files, download dependencies and compile executable
FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./adguardian

# Run stage - Copy executable from build stage and run it
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/adguardian ./
CMD [ "./adguardian" ]
