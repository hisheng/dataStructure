package crypto_hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/go-playground/assert/v2"
	"io"
	"testing"
)

//密码散列函数（英语：Cryptographic hash function），又译为加密散列函数
//加密

func TestSha1(t *testing.T) {
	//sha1 则返回值是一个 40 字符长度的十六进制数字
	message := "hello world hisheng"
	h := sha1.New()
	io.WriteString(h, message)
	t.Logf("% x", h.Sum(nil))
	t.Log(h.Sum(nil))
	s := fmt.Sprintf("%x", h.Sum(nil))
	t.Log(s)
	assert.Equal(t, s, sha1Hash(message))

	message2 := "hi hisheng hello"
	io.WriteString(h, message2)
	t.Logf("% x", h.Sum(nil))
	s1 := fmt.Sprintf("%x", h.Sum(nil))

	assert.Equal(t, s1, sha1Hash(message+message2))

}

func sha1Hash(message string) string {
	h := sha1.New()
	io.WriteString(h, message)
	s := fmt.Sprintf("%x", h.Sum(nil))
	return s
}

func sha256Hash(message string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(message)))
}

func sha256Hash2(message string) string {
	h := sha256.New()
	io.WriteString(h, message)
	s := fmt.Sprintf("%x", h.Sum(nil))
	return s
}

func TestSha256(t *testing.T) {
	message := "hello world hisheng"
	assert.Equal(t, sha256Hash(message), sha256Hash2(message))
}
