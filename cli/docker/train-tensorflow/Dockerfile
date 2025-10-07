FROM trustmint/train-base:latest

# Install TensorFlow
RUN pip install --no-cache-dir tensorflow pandas numpy

# Copy training script
COPY train.py /workspace/train.py

ENTRYPOINT ["python", "/workspace/train.py"]
