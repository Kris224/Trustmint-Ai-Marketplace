FROM trustmint/train-base:latest

# Install scikit-learn
RUN pip install --no-cache-dir scikit-learn pandas numpy joblib

# Copy training script
COPY train.py /workspace/train.py

ENTRYPOINT ["python", "/workspace/train.py"]
