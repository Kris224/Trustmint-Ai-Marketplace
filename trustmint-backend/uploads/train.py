import os
import time

print("--- Python training script started ---")

dataset_path = './dataset/data.csv'
model_output_path = './output/model.pkl'

try:
    with open(dataset_path, 'r') as f:
        lines = f.readlines()
        print(f"   - Successfully read {len(lines)} lines from the dataset.")
except Exception as e:
    print(f"   - Error reading dataset: {e}")
    exit(1)

print("   - Simulating model training for 5 seconds...")
# time.sleep(5) 

try:
    with open(model_output_path, 'w') as f:
        f.write("This is a dummy trained model.")
    print(f"   - Model saved successfully to {model_output_path}")
except Exception as e:
    print(f"   - Error saving model: {e}")
    exit(1)

print("--- Python training script finished ---")
