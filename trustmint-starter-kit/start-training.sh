
set -e

DOCKER_IMAGE="20demo/trustmint-cli:latest"
CONFIG_FILE="trustmint.yml"

echo "▶️  Starting Trustmint Workflow..."
echo "----------------------------------------------------"
echo "1/2: Verifying setup and parsing configuration..."

if ! command -v yq &> /dev/null; then
    echo "❌ Error: 'yq' is not installed. Please install it: sudo snap install yq"
    exit 1
fi
if ! docker info > /dev/null 2>&1; then echo "❌ Error: Docker is not running." && exit 1; fi
if [ ! -f "$CONFIG_FILE" ]; then echo "❌ Error: '$CONFIG_FILE' not found." && exit 1; fi

CONFIG_HASH=$(sha256sum "$CONFIG_FILE" | awk '{ print $1 }')
echo "🔑 Config Hash Generated: ${CONFIG_HASH:0:12}..."

NETWORK_MODE=$(yq e '.environment.network' $CONFIG_FILE)
READ_ONLY_ROOT=$(yq e '.environment.read_only_root' $CONFIG_FILE)

DOCKER_FLAGS="-it --rm --network $NETWORK_MODE"
if [ "$READ_ONLY_ROOT" = "true" ]; then
    DOCKER_FLAGS="$DOCKER_FLAGS --read-only"
fi

VOLUME_COUNT=$(yq e '.environment.volumes | length' $CONFIG_FILE)
for i in $(seq 0 $(($VOLUME_COUNT - 1))); do
    SOURCE=$(yq e ".environment.volumes[$i].source" $CONFIG_FILE)
    TARGET=$(yq e ".environment.volumes[$i].target" $CONFIG_FILE)
    READ_ONLY_VOL=$(yq e ".environment.volumes[$i].read_only" $CONFIG_FILE)
    
    ABS_SOURCE="$(pwd)/$SOURCE"
    
    VOLUME_FLAG="-v $ABS_SOURCE:$TARGET"
    if [ "$READ_ONLY_VOL" = "true" ]; then
        VOLUME_FLAG="$VOLUME_FLAG:ro"
    fi
    DOCKER_FLAGS="$DOCKER_FLAGS $VOLUME_FLAG"
done

echo "⚙️  Applying all settings from $CONFIG_FILE..."
echo "🚀 Pulling latest runner from Docker Hub..."
docker pull $DOCKER_IMAGE > /dev/null
echo "✅ System ready."

echo "2/2: Launching container..."
echo "----------------------------------------------------"
echo "✅ Success! You are now inside the secure sandbox."
echo "   Next, run 'trustmint train', then type 'exit'."
echo "----------------------------------------------------"

docker run \
  $DOCKER_FLAGS \
  -e CONFIG_HASH="$CONFIG_HASH" \
  $DOCKER_IMAGE