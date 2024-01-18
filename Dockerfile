FROM golang:1.21.3-alpine
WORKDIR /go/src/app

# Install go tools
RUN go install github.com/go-task/task/v3/cmd/task@v3.30.1

# Copy and install go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the app
RUN task build

# Run the app
CMD ["task", "serve"]
