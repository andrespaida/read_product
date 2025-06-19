# Dockerfile

FROM golang:1.21

# Create working directory
WORKDIR /app

# Copy dependency files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o main .

# Expose the port
EXPOSE 4001

# Run the app
CMD ["./main"]