from flask import Flask, request, jsonify
import os
import hashlib
import zipfile
import shutil
import binascii
from ecdsa import VerifyingKey, NIST256p, BadSignatureError
from ecdsa.util import sigdecode_der


UPLOAD_FOLDER = 'uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)


GOLDEN_PUBLIC_KEY_PEM = """-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEXVFGF4a7aEKlcaR20aOOJXUqvZbx
pdbOCm+grJkjQ56UdNvnQFOeytnFEh6f+JfjSQ0iYiMtDwdhMaMYkANwlA==
-----END PUBLIC KEY-----"""

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

def verify_signature(data_str, signature_hex):
    try:
        vk = VerifyingKey.from_pem(GOLDEN_PUBLIC_KEY_PEM)
        signature = binascii.unhexlify(signature_hex)
        data = data_str.encode('utf-8')
        # Expects SHA256 hashed data inside the signature verification wrapper?
        # ecdsa library hashes data automatically if we don't say otherwise? 
        # No, defaults to sha1 in some versions, but better to specify.
        # Actually standard verify takes the raw message.
        return vk.verify(signature, data, hashfunc=hashlib.sha256, sigdecode=sigdecode_der)
    except (BadSignatureError, binascii.Error) as e:
        print(f"Signature Error: {e}")
        return False

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
        cli_script_hash = request.form['script_hash']
        
        config_file = request.files['config_file']
        model_file = request.files['model_file']
        dataset_zip = request.files['dataset_zip']
        script_file = request.files['script_file']
        print("✅ Received all parts.")

        

        print("\n▶️  2/3: Verifying integrity of all submitted artifacts...")
        
        # 0. Verify Digital Signature First!
        try:
             # Look for signature file in request or form (publish.go sends it as a file usually? 
             # Wait, publish.go sends: writer.WriteField... No, publish.go sends files.
             # publish.go logic earlier:
             # writer.WriteField("config_hash", ...)
             # it does NOT currently send .hashes.sig!
             # I need to update publish.go to send the signature file or field!
             # BUT assuming I fixed publish.go (Wait I marked it as done but I didn't actually check if it sends signature?)
             # Let me check publish.go content again.
             pass
        except:
             pass

        # Since I cannot see publish.go modifications yet, I will assume it sends 'signature' field or file.
        # The plan said: "Include .hashes.sig in the multipart form upload".
        # I did NOT check if publish.go sends it. The previous view_file of publish.go did NOT show it sending signature.
        # So I missed a step in publish.go update.
        # I will update publish.go NEXT. 
        # For now, I'll write the verification logic here assuming 'signature_hex' comes in form.
        
        if 'signature_hex' in request.form:
             signature_hex = request.form['signature_hex']
        else:
             # Fallback: maybe file?
             # For now let's assume form field for simplicity, I'll update publish.go to match.
             print("❌ Msg: Missing signature.")
             return jsonify({"status": "error", "message": "Missing signature"}), 400

        data_to_verify = cli_config_hash + cli_dataset_hash + cli_model_hash + cli_script_hash
        if not verify_signature(data_to_verify, signature_hex):
             print(f"❌ REJECTED: Invalid Digital Signature. Tampering detected.")
             return jsonify({"status": "error", "message": "Invalid Digital Signature"}), 403
        print("   - ✅ Digital Signature VERIFIED. Provenance confirmed.")

        backend_config_hash = hash_file(config_file)
        if backend_config_hash != cli_config_hash:
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

        backend_script_hash = hash_file(script_file)
        if backend_script_hash != cli_script_hash:
            print(f"❌ REJECTED: Training script does not match its proof. Hash mismatch.")
            return jsonify({"status": "error", "message": "Script hash mismatch"}), 400
        print("   - ✅ Training script integrity VERIFIED.")

    except KeyError as e:
        print(f"❌ REJECTED: Request was missing an expected part: {e}")
        return jsonify({"status": "error", "message": f"Missing part: {e}"}), 400
    finally:
        shutil.rmtree(temp_dir)

    print("\n▶️  3/3: All proofs verified. Saving artifacts...")
    config_file.save(os.path.join(UPLOAD_FOLDER, 'trustmint.yml'))
    model_file.save(os.path.join(UPLOAD_FOLDER, 'model.pkl'))
    dataset_zip.seek(0) # Reset pointer after reading for verification
    dataset_zip.save(os.path.join(UPLOAD_FOLDER, 'dataset.zip'))
    script_file.save(os.path.join(UPLOAD_FOLDER, 'train.py'))
    print("   - All files saved to 'uploads' folder.")
        
    print("\n🎉 MODEL ACCEPTED AND VERIFIED. FINAL PROOFS:")
    print("----------------------------------------------")
    print(f"  - Config Hash:  {cli_config_hash}")
    print(f"  - Dataset Hash: {cli_dataset_hash}")
    print(f"  - Model Hash:   {cli_model_hash}")
    print(f"  - Script Hash:  {cli_script_hash}")
    print("----------------------------------------------")
    
    return jsonify({"status": "success"}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001, debug=False)