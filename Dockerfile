FROM python:3.10-slim

# Install Go and required build tools
RUN apt-get update && \
    apt-get install -y wget gcc git && \
    wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz && \
    rm go1.21.0.linux-amd64.tar.gz && \
    apt-get clean

# Add Go to PATH
ENV PATH=$PATH:/usr/local/go/bin

# Set the working directory to the project root
WORKDIR /app

# Copy the entire project (backend, CLI, and blockchain artifacts)
COPY . /app

# Switch to the Go CLI directory
WORKDIR /app/trustmint-cli
RUN go mod download

# Precompile the CLI with a 42-character dummy wallet address
RUN go build -ldflags "-X trustmint.com/cli/cmd.WalletAddress=0x0000000000000000000000000000000000000000 -X trustmint.com/cli/cmd.BackendURL=https://trustmint-ai-marketplace.onrender.com" -o /app/trustmint-backend/trustmint-template .

# Set working directory to the backend
WORKDIR /app/trustmint-backend

# Install python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Run gunicorn when the container launches. 
# Gunicorn will automatically use the $PORT variable provided by Render.
CMD ["gunicorn", "--timeout", "120", "app:app"]
