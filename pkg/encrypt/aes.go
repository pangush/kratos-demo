package encrypt

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上，根据key来决定填充blocksize
*/

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)

	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes加密，填充模式由key决定，16位，24,32分别对应AES-128, AES-192, or AES-256.源码好像是写死16了

func Encrypt(rawData, key []byte) (string, error) {
	data, err := ecbEncrypt(rawData, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(rawData string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := ecbDecrypt(data, key)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

func ecbDecrypt(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted), nil
}

func ecbEncrypt(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}
