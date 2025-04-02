package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func EncAES(data []byte, key []byte) ([]byte, error) {
	hash, _ := GetMD5(data)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, hash)
	encBuffer := make([]byte, len(data))
	stream.XORKeyStream(encBuffer, data)
	return append(hash, encBuffer...), nil
}

func DecAES(data []byte, key []byte) ([]byte, error) {
	// MD5[16 bytes] + Data[n bytes]
	dataLen := len(data)
	if dataLen <= 16 {
		return nil, ErrEntityInvalid
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, data[:16])
	decBuffer := make([]byte, dataLen-16)
	stream.XORKeyStream(decBuffer, data[16:])
	hash, _ := GetMD5(decBuffer)
	if !bytes.Equal(hash, data[:16]) {
		return nil, ErrFailedVerification
	}
	return decBuffer[:dataLen-16], nil
}
