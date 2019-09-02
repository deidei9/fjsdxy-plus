package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//利用MD5加密字符串
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//对字符串进行SHA1哈希
func GetSH1(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串base64加密
func GetBase64(str string) ([]byte, error) {
	encodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return encodeString, nil
}

//微信用户数据AES解密 =>占位字符原因废弃，使用github.com/xlstudio/wxbizdatacrypt
func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil, errors.New("解密失败")
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	//获取的数据尾端有'/x0f'占位符,去除它
	for i, ch := range origData {
		if ch == '\x10' {
			origData[i] = ' '
		}
		if ch == '\x0f' {
			origData[i] = ' '
		}
	}
	return origData, nil
}
