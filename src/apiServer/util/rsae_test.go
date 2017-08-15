package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * base64 加密
 */
func TestBase64Encode(t *testing.T) {
	// assert.Len(t, GetString(2), 2)
	// fmt.Println(GetString(20))
	r := NewRsae().Base64Encode([]byte("diogoxiang"))
	// ZGlvZ294aWFuZw==
	fmt.Println(r)

}

/**
 * base64 解密
 */
func TestBase64Decode(t *testing.T) {
	r, _ := NewRsae().Base64Decode("ZGlvZ294aWFuZw==")
	assert.Equal(t, string(r), "diogoxiang")
}

func TestM516(t *testing.T) {
	r := NewRsae().Md516("diogoxiang")
	// edb0ebc639387365
	fmt.Println(r)
	assert.Equal(t, string(r), "edb0ebc639387365")
}

func TestRsaKey(t *testing.T) {
	RsaFile()
}

func TestRsa(t *testing.T) {
	msg := "diogoxiang"
	publicPath := "public.pem"
	publicKey, _ := NewFile().Read(publicPath)

	privatePath := "private.pem"
	privateKey, _ := NewFile().Read(privatePath)

	pubKey := ParsePublicKey(publicKey)
	if pubKey == nil {
		fmt.Println("Parse PublicKey err")
	}

	// 加密数
	// encryptPKCS15 := tool.EncryptPKCS1v15(pubKey, msg)
	// if encryptPKCS15 == nil {
	// 	fmt.Println("rsa EncryptPKCS1v15 err")
	// }
	// fmt.Println("EncryptPKCS1v15 string:" + string(encryptPKCS15) + "\n")

	encryptOAEP := EncryptOAEP(pubKey, []byte(msg))
	if encryptOAEP == nil {
		fmt.Println("rsa EncryptOAEP err")
	}
	fmt.Println("EncryptOAEP string :" + string(encryptOAEP) + "\n")
	file := "key.txt"
	err := NewFile().Write(file, encryptOAEP)
	if err != nil {
		fmt.Println(err)
	}

	// 解析出私钥
	priKey := ParsePrivateKey(privateKey)
	if priKey == nil {
		fmt.Println("Parse PrivateKey err")
	}

	decryptOAEP := DecryptOAEP(priKey, encryptOAEP)
	if decryptOAEP == nil {
		fmt.Println("rsa DecryptOAEP err")
	}
	fmt.Println("DecryptOAEP string:" + string(decryptOAEP) + "\n")

}
