// Package aes 提供对标准库 crypto/aes 的封装,简化对称加密和解密的操作
package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

const (
	keySize = 32
)

type Encryptor struct {
	key  []byte
	mode cipher.Block
}

type Option func(e *Encryptor)

func generateKeyFromSalt(salt string, size int) []byte {
	key := make([]byte, size)
	copy(key, salt)
	return key
}

func NewEncryptor(salt string, opts ...Option) (*Encryptor, error) {
	e := &Encryptor{
		key: generateKeyFromSalt(salt, keySize),
	}

	// 创建加密器
	cipherBlock, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	e.mode = cipherBlock

	for _, opt := range opts {
		opt(e)
	}
	return e, nil
}

func (e *Encryptor) Encrypt(plainText []byte) (encrypted []byte, err error) {
	// 处理待加密内容，使其长度为分组长度的整数倍
	// 需要填充多少长度
	padding := aes.BlockSize - len(plainText)%aes.BlockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)

	// 填充
	plainText = append(plainText, pad...)

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv) // 每次加密生成随机初始向量
	if err != nil {
		return nil, err
	}

	// 对与待加密数据，会按照16字节的块大小独立进行加密（结合前面的块）
	// 因为第一个块没有前一个块，所以需要一个初始向量iv，iv由随机数生成

	// 头部预留给初始向量
	// 头部填充初始向量
	c := cipher.NewCBCEncrypter(e.mode, iv)
	result := make([]byte, len(plainText)+aes.BlockSize)
	c.CryptBlocks(result[aes.BlockSize:], plainText)
	copy(result[:aes.BlockSize], iv)

	return result, nil
}

func (e *Encryptor) Decrypt() {

}
