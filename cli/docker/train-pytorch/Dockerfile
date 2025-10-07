FROM trustmint/train-base:latest

# Install PyTorch (CPU-only)
RUN pip install --no-cache-dir torch torchvision pandas numpy

# Copy training script
COPY train.py /workspace/train.py

ENTRYPOINT ["python", "/workspace/train.py"]
