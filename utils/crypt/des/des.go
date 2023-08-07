package des

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

type Encryption struct {
	Key []byte
}

// NewDesCrypt 创建一个des加密对象 传入密钥 eg:var pwDeKey = []byte{3, 2, 3, 2, 5, 5, 1, 1}  加密解密 key 必须8位字节-----DES算法
func NewDesCrypt(pwDeKey []byte) *Encryption {
	return &Encryption{Key: pwDeKey}
}
func (e *Encryption) Init(pwDeKey []byte) {
	e.Key = pwDeKey
}

func (e *Encryption) Encrypt(text string) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(e.Key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func (e *Encryption) Decrypt(decrypted string) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(e.Key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
