// internal/crypto/crypto.go
package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
)

// FileSHA256Hex computes SHA-256 hash of a file and returns hex string.
func FileSHA256Hex(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// MerkleRoot computes a simple Merkle root from a list of hex-encoded hashes.
// It expects input hashes as hex strings. The algorithm concatenates pairs of hashes,
// hashes the concatenation, and repeats until a single root remains.
func MerkleRoot(hashes []string) string {
	if len(hashes) == 0 {
		return ""
	}
	// copy to avoid mutating input
	h := make([]string, len(hashes))
	copy(h, hashes)

	// ensure deterministic ordering
	sort.Strings(h)

	for len(h) > 1 {
		var next []string
		for i := 0; i < len(h); i += 2 {
			if i+1 < len(h) {
				combined := h[i] + h[i+1]
				sum := sha256.Sum256([]byte(combined))
				next = append(next, hex.EncodeToString(sum[:]))
			} else {
				// odd one out, carry forward
				next = append(next, h[i])
			}
		}
		h = next
	}
	return h[0]
}

// MerkleRootFromFiles computes SHA256 for each provided file path and returns the Merkle root.
func MerkleRootFromFiles(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("no paths provided")
	}
	// deterministic order
	sort.Strings(paths)
	var hashes []string
	for _, p := range paths {
		hv, err := FileSHA256Hex(p)
		if err != nil {
			return "", err
		}
		hashes = append(hashes, hv)
	}
	return MerkleRoot(hashes), nil
}


// HashFile is retained for backward compatibility with earlier code.
func HashFile(path string) (string, error) {
    return FileSHA256Hex(path)
}
