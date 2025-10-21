set -e

DOCKER_IMAGE="kris20204/trustmint-runner:latest"

echo "▶️  Starting Trustmint ONLINE Publishing..."
echo "----------------------------------------------------"
echo "1/2: Verifying that proofs exist..."
if [ ! -f ".hashes.yaml" ]; then
    echo "❌ Error: Proofs file (.hashes.yaml) not found."
    echo "   Please run './start-training.sh' and the 'trustmint train' command first."
    exit 1
fi
echo "✅ Proofs found."

echo "2/2: Launching ONLINE container to publish..."

docker run --rm \
  --network="host" \
  -v "$(pwd):/app/project" \
  $DOCKER_IMAGE \
  bash -c "trustmint publish"

echo "----------------------------------------------------"
echo "🎉 Workflow Finished! Model has been published for verification."