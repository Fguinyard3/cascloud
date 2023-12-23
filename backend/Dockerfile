# Build Stage
FROM golang:alpine AS builder

# Install any dependencies your project may need
RUN apk add --no-cache git

# Set the working directory
WORKDIR /go/src/app

# Copy the entire project into the container
COPY . .

# Fetch dependencies and build the executable
RUN go get -d -v ./...
RUN go build -o /go/bin/app

# Final Stage
FROM alpine:latest

# Copy only the executable from the build stage
COPY --from=builder /go/bin/app /app

# Expose the port your application will run on
EXPOSE 8000

# Command to run the executable
CMD ["/app"]

