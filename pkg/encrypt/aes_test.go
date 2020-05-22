package encrypt

import (
	"kratos-demo/internal/conf"
	"testing"
)

func TestEncrypt(t *testing.T) {
	cfg := conf.Conf
	raw := []byte("13260168343")
	key := []byte(cfg.Aes.AesKey)
	str, err := Encrypt(raw, key)
	if err == nil {
		t.Log("suc", str)
	} else {
		t.Fatal("fail", err)
	}
}

func TestDncrypt(t *testing.T) {
	raw := "uymMptu+dnYKml4LrX/8SA=="
	key := []byte(conf.Conf.Aes.AesKey)
	str, err := Decrypt(raw, key)
	if err == nil {
		t.Log("suc", str)
	} else {
		t.Fatal("fail", err)
	}
}
