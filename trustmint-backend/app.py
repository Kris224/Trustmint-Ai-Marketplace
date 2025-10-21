from flask import Flask, request, jsonify
import os
import hashlib
import zipfile
import shutil


UPLOAD_FOLDER = 'uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)


GOLDEN_CONFIG_HASH = "3ae6951d92f9052c77e15a33a099afcfc25a7d5c509767a3b50aff9173b8ae48"

app = Flask(__name__)



def hash_directory(path):
    """Calculates a single SHA256 hash for all files in a directory."""
    hasher = hashlib.sha256()
    for root, dirs, files in sorted(os.walk(path)):
        for name in sorted(files):
            file_path = os.path.join(root, name)
            with open(file_path, 'rb') as f:
                while chunk := f.read(4096):
                    hasher.update(chunk)
    return hasher.hexdigest()

def hash_file(file_storage):
    """Calculates the SHA256 hash of a file received from the CLI."""
    hasher = hashlib.sha256()
    file_storage.seek(0)  # Go to the beginning of the file
    while chunk := file_storage.read(4096):
        hasher.update(chunk)
    file_storage.seek(0)  # Rewind the file in case we need to save it
    return hasher.hexdigest()

@app.route('/publish', methods=['POST'])
def handle_publish():
    print("\n\n==============================================")
    print("--- ✅ INCOMING PUBLISH REQUEST ---")
    print("==============================================")
    
    temp_dir = os.path.join(UPLOAD_FOLDER, "temp_verification")
    os.makedirs(temp_dir, exist_ok=True)

    try:
    
        print("▶️  1/3: Receiving proofs and files from the CLI...")
        cli_config_hash = request.form['config_hash']
        cli_dataset_hash = request.form['dataset_hash']
        cli_model_hash = request.form['model_hash']
        
        config_file = request.files['config_file']
        model_file = request.files['model_file']
        dataset_zip = request.files['dataset_zip']
        print("✅ Received all parts.")

        

        print("\n▶️  2/3: Verifying integrity of all submitted artifacts...")

        backend_config_hash = hash_file(config_file)
        if backend_config_hash != GOLDEN_CONFIG_HASH or backend_config_hash != cli_config_hash:
            print(f"❌ REJECTED: Config file was altered. Hash mismatch.")
            return jsonify({"status": "error", "message": "Config hash mismatch"}), 400
        print("   - ✅ Config integrity VERIFIED.")

        backend_model_hash = hash_file(model_file)
        if backend_model_hash != cli_model_hash:
            print(f"❌ REJECTED: Model file does not match its proof. Hash mismatch.")
            return jsonify({"status": "error", "message": "Model hash mismatch"}), 400
        print("   - ✅ Model integrity VERIFIED.")

        dataset_zip_path = os.path.join(temp_dir, 'dataset.zip')
        dataset_unzip_path = os.path.join(temp_dir, 'dataset')
        dataset_zip.save(dataset_zip_path)
        with zipfile.ZipFile(dataset_zip_path, 'r') as zip_ref:
            zip_ref.extractall(dataset_unzip_path)
        
        backend_dataset_hash = hash_directory(dataset_unzip_path)
        if backend_dataset_hash != cli_dataset_hash:
            print(f"❌ REJECTED: Dataset does not match its proof. Hash mismatch.")
            return jsonify({"status": "error", "message": "Dataset hash mismatch"}), 400
        print("   - ✅ Dataset integrity VERIFIED.")

    except KeyError as e:
        print(f"❌ REJECTED: Request was missing an expected part: {e}")
        return jsonify({"status": "error", "message": f"Missing part: {e}"}), 400
    finally:
        shutil.rmtree(temp_dir)

    print("\n▶️  3/3: All proofs verified. Saving artifacts...")
    config_file.save(os.path.join(UPLOAD_FOLDER, 'trustmint.yml'))
    model_file.save(os.path.join(UPLOAD_FOLDER, 'model.pkl'))
    dataset_zip.save(os.path.join(UPLOAD_FOLDER, 'dataset.zip'))
    print("   - All files saved to 'uploads' folder.")
        
    print("\n🎉 MODEL ACCEPTED AND VERIFIED. FINAL PROOFS:")
    print("----------------------------------------------")
    print(f"  - Config Hash:  {cli_config_hash}")
    print(f"  - Dataset Hash: {cli_dataset_hash}")
    print(f"  - Model Hash:   {cli_model_hash}")
    print("----------------------------------------------")
    
    return jsonify({"status": "success"}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=False)