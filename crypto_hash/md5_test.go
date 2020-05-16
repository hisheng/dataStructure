package crypto_hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"testing"
)

func TestMd5(t *testing.T) {
	//方法1
	m5 := md5.New()
	io.WriteString(m5, "mynameishishneg")
	m5s := fmt.Sprintf("%x", m5.Sum(nil))
	t.Log(m5s)

	m5 = md5.New()
	io.WriteString(m5, "hiword")
	m5s = fmt.Sprintf("%x", m5.Sum(nil))
	t.Log(m5s)
	//方法2
	t.Log(fmt.Sprintf("%x", md5.Sum([]byte("mynameishishneg"))))
	t.Log(fmt.Sprintf("%x", md5.Sum([]byte("hiword"))))

	//方法3
	m5 = md5.New()
	m5.Write([]byte("mynameishishneg"))
	m5str := m5.Sum(nil)
	t.Log(fmt.Sprintf("%x", m5str))
}
