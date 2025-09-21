# Stage 1: The Build Stage
# Use a full Go image to compile the application.
FROM golang:1.25 AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies.
# This step is cached by Docker, speeding up subsequent builds.
COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the entire source code into the container.
COPY . .

# Build the Go application for a Linux environment.
# CGO_ENABLED=0 creates a statically linked binary, which is a best practice for small images.
# The `-o` flag specifies the output name and path of the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /simplemath ./cmd/main.go

# ---

# Stage 2: The Run Stage
# Use a minimal, lightweight image to run the application.
# Alpine is a great choice as it's very small.
FROM alpine:latest

# Set the working directory.
WORKDIR /root/

# Copy the built binary from the `builder` stage.
COPY --from=builder /simplemath .

# Copy the static files into the container.
# This is crucial for serving your HTML, CSS, and favicon.
# The files will be located at /root/statics in the final image.
COPY --from=builder /app/statics ./statics

# Expose the port your application listens on.
# Your app's code suggests port 8080 is likely used.
EXPOSE 8080

# Command to run the application.
CMD ["./simplemath"]