package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type AesCBC struct {
	key []byte
}

func New(key string) (ac *AesCBC) {
	ac = &AesCBC{}
	ac.key = ac.string2bytes(key)
	ac.key = ac.generateSha256Key()
	return
}

func (ac *AesCBC) generateSha256Key() (res []byte) {
	h := sha256.New()
	h.Write(ac.key)
	res = h.Sum(nil)
	return
}

func (ac *AesCBC) string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func (ac *AesCBC) pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func (ac *AesCBC) pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (ac *AesCBC) Encrypt(rawDataStr string) ([]byte, error) {
	if rawDataStr == "" {
		return nil, errors.New("aes Encrypt invalid rawDataStr")
	}
	data := ac.string2bytes(rawDataStr)
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data = ac.pKCS7Padding(data, blockSize)
	mode := cipher.NewCBCEncrypter(block, ac.key[:blockSize])
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	return ciphertext, nil
}

func (ac *AesCBC) Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, ac.key[:bs])
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	data = ac.pKCS7UnPadding(data)
	return data, nil
}

func (ac *AesCBC) Base64Encode(data []byte) (res string) {
	return base64.StdEncoding.EncodeToString(data)
}

func (ac *AesCBC) Base64Decode(msg string) (res []byte, err error) {
	return base64.StdEncoding.DecodeString(msg)
}

const (
	key  = "2923e9cb8d967d08f5878ca0356eb3ca"
	data = `{"name":"张大而","sex":"男","wechat":"aabb","qq":"3333244411","phone":"14578780998","email":"cccc@test.com","pics":{"prove":"http://uat-i0.hdslb.com/bfs/archive/5b1cf41dbd02d1f5f127c3d64eafdd3a22c82afd.jpg_60x60.jpg","positive_side":"http://uat-i0.hdslb.com/bfs/archive/9b9859e3555d0e54c6f16ce7fc13f80d13bc2052.jpg_60x60.jpg","negative_side":"http://uat-i0.hdslb.com/bfs/archive/5a78e855971b054b43e88ec05cc3c5f054462905.jpg_60x60.jpg"}}`
)

func main() {
	myAes := New(key)
	crypted, err := myAes.Encrypt(data)
	if err != nil {
		panic(err)
	}

	encoded := myAes.Base64Encode(crypted)
	fmt.Println("占用空间大小:", unsafe.Sizeof(encoded), "长度:", len(encoded))
	fmt.Println("base64编码:", encoded)

	decoded, err := myAes.Base64Decode(encoded)
	if err != nil {
		fmt.Println("Base64Decode:", err)
		return
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println("base64解码:", string(decoded))
	fmt.Println("------------------------------------------")

	data, err := myAes.Decrypt(decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
