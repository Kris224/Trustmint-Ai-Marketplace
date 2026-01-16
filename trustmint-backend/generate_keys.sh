#!/bin/bash

# 1. Generate Private Key (Elliptic Curve - Prime256v1)
openssl ecparam -name prime256v1 -genkey -noout -out private.pem

# 2. Extract Public Key form Private Key
openssl ec -in private.pem -pubout -out public.pem

echo "✅ Keys generated successfully!"
echo "   - private.pem (KEEP SECRET! Used by CLI)"
echo "   - public.pem  (Public. Used by Backend)"
