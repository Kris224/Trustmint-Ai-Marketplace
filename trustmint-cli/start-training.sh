#!/bin/bash
set -e

# --- Configuration ---
DOCKER_IMAGE="20demo/trustmint-cli:latest"
CONFIG_FILE="trustmint.yml"

# --- Pre-flight Checks ---
echo "✅ Verifying project setup..."
# ... (all checks for docker, yq, files, etc. remain the same) ...
if ! docker info > /dev/null 2>&1; then echo "❌ Error: Docker is not running." && exit 1; fi
if [ ! -f "$CONFIG_FILE" ]; then echo "❌ Error: '$CONFIG_FILE' not found." && exit 1; fi
if [ ! -d "dataset" ]; then echo "❌ Error: 'dataset' folder not found." && exit 1; fi

# --- Hashing ---
echo "🔑 Calculating hash of the configuration file..."
CONFIG_HASH=$(sha256sum "$CONFIG_FILE" | awk '{ print $1 }')
echo "   - Config Hash: $CONFIG_HASH"

echo "🚀 Pulling the latest secure runner image..."
docker pull $DOCKER_IMAGE

# =================================================================
# STAGE 1: SECURE OFFLINE TRAINING
# =================================================================
echo "🔒 Launching OFFLINE container for secure training..."

docker run --rm \
  --network none \
  --read-only \
  -e CONFIG_HASH="$CONFIG_HASH" \
  -v "$(pwd)/dataset:/app/dataset:ro" \
  -v "$(pwd):/app/project" \
  $DOCKER_IMAGE \
  train # This tells the CLI to run the 'train' command

echo "✅ Training complete. Proofs saved to .hashes.yaml."

# =================================================================
# STAGE 2: ONLINE PUBLISHING
# =================================================================
echo "🚀 Launching ONLINE container for publishing..."

docker run --rm \
  --network="host" \
  -v "$(pwd):/app/project" \
  $DOCKER_IMAGE \
  publish # This tells the CLI to run the 'publish' command

echo "🎉 Process complete! Model has been published."