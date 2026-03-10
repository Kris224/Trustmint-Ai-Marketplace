"""
Sample training script for Trustmint testing.
Reads from dataset/data.csv and trains a simple linear regression model.
Produces a unique model.pkl each run based on the data.
"""

import os
import csv
import pickle
import time

DATASET_DIR = "./dataset"
OUTPUT_PATH = "./output/model.pkl"
os.makedirs("./output", exist_ok=True)

print("--- Python training script started ---")

# Read all CSV files from the dataset directory
rows = []
for fname in os.listdir(DATASET_DIR):
    if fname.endswith(".csv"):
        with open(os.path.join(DATASET_DIR, fname), newline="") as f:
            reader = csv.DictReader(f)
            for row in reader:
                rows.append(row)

print(f"   - Loaded {len(rows)} rows from dataset.")

if not rows:
    raise ValueError("Dataset is empty! Add data to dataset/data.csv")

# Extract numeric columns for simple linear regression
headers = list(rows[0].keys())
numeric_cols = []
for h in headers:
    try:
        float(rows[0][h])
        numeric_cols.append(h)
    except (ValueError, TypeError):
        pass

if len(numeric_cols) < 2:
    # Not enough numeric cols — store raw stats only
    feature_col, label_col = headers[0], headers[-1]
    values = [r[feature_col] for r in rows]
    model = {
        "type":        "stats",
        "feature_col": feature_col,
        "label_col":   label_col,
        "n_samples":   len(rows),
        "values":      values,
        "trained_at":  time.time(),
    }
else:
    feature_col = numeric_cols[0]
    label_col   = numeric_cols[-1]

    X = [float(r[feature_col]) for r in rows]
    y = [float(r[label_col])   for r in rows]

    n     = len(X)
    x_bar = sum(X) / n
    y_bar = sum(y) / n

    # Simple linear regression: y = slope * x + intercept
    numerator   = sum((X[i] - x_bar) * (y[i] - y_bar) for i in range(n))
    denominator = sum((X[i] - x_bar) ** 2 for i in range(n)) or 1e-9
    slope       = numerator / denominator
    intercept   = y_bar - slope * x_bar

    # Predictions and R² score
    preds = [slope * x + intercept for x in X]
    ss_res = sum((y[i] - preds[i]) ** 2 for i in range(n))
    ss_tot = sum((y[i] - y_bar)    ** 2 for i in range(n)) or 1e-9
    r2 = 1 - ss_res / ss_tot

    print(f"   - Feature: '{feature_col}', Label: '{label_col}'")
    print(f"   - Slope: {slope:.4f}, Intercept: {intercept:.4f}, R²: {r2:.4f}")

    model = {
        "type":        "linear_regression",
        "feature_col": feature_col,
        "label_col":   label_col,
        "slope":       slope,
        "intercept":   intercept,
        "r2":          r2,
        "n_samples":   n,
        "trained_at":  time.time(),  # ensures unique hash each run
    }

with open(OUTPUT_PATH, "wb") as f:
    pickle.dump(model, f)

print(f"   - Model saved to {OUTPUT_PATH}")
print("--- Python training script finished ---")
