# --- Stage 1: The "Builder" ---
FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# This builds a self-contained executable.
ARG PRIVATE_KEY
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X 'trustmint.com/cli/cmd.AppPrivateKey=${PRIVATE_KEY}'" -o /trustmint .

# --- Stage 2: The Final Image ---
FROM pytorch/pytorch:2.1.0-cuda12.1-cudnn8-runtime
WORKDIR /app/project
# Copy the compiled program from the 'builder' stage
COPY --from=builder /trustmint /usr/local/bin/trustmint
RUN chmod +x /usr/local/bin/trustmint

# start an interactive bash shell, which will wait for your commands.
CMD ["bash"]