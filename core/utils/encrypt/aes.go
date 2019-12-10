package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	errors2 "errors"

	"github.com/treeyh/soc-go-common/core/errors"
)

func AesDecrypt(key string, encrypt string) (string, errors.AppError) {
	kbs := SHA256(key)
	decode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.EncryptDecryptFail, err)
	}
	if len(decode) < aes.BlockSize {
		return "", errors.NewAppErrorByExistError(errors.EncryptDecryptFail, errors2.New("密文太短啦"))
	}
	iv := decode[:aes.BlockSize]
	block, err := aes.NewCipher(kbs)
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.EncryptDecryptFail, err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(decode))
	blockMode.CryptBlocks(plantText, decode)
	plantText = PKCS7UnPadding(plantText)
	plantText = plantText[aes.BlockSize:]
	return string(plantText), nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func SHA256(source string) []byte {
	mac := sha256.New()
	mac.Write([]byte(source))
	return mac.Sum(nil)
}
