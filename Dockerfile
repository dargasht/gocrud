# Stage 1: Build the Go binary
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# # Download Go modules
RUN go mod download

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 go build -o app ./cmd/api/.

# Stage 2: Create a lightweight image with only the binary
FROM gcr.io/distroless/base-debian12

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/app .

# Copy whatever else you need here

# Expose the port the app will run on
EXPOSE 42069

# Command to run the binary
CMD ["./app"]