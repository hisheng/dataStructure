package bloom_filter

import (
	"crypto/sha256"
	"testing"
)

func TestHashData(t *testing.T) {
	sha_data := sha256.Sum256([]byte{'h', 'e', 'l', 'l', 'o'})
	t.Log(sha_data)
	t.Log(HashData([]byte{'h', 'e', 'l', 'l', 'o'}, 0)) //3182187045311309425 19位
}
