import os
import joblib
from sklearn.datasets import load_iris
from sklearn.linear_model import LogisticRegression

# Load dataset (for demo)
X, y = load_iris(return_X_y=True)

# Train model
model = LogisticRegression(max_iter=200)
model.fit(X, y)

# Save output
os.makedirs("/workspace/output", exist_ok=True)
joblib.dump(model, "/workspace/output/model.pkl")

print("✅ Training complete. Model saved to /workspace/output/model.pkl")
