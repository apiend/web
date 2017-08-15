package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

// Rsae rsae
type Rsae struct{}

// NewRsae new rsae
func NewRsae() *Rsae {
	return &Rsae{}
}

// Base64Encode base64 encode
func (r Rsae) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//Base64Decode base64 descode
func (r Rsae) Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// Md516 md5 16
func (r Rsae) Md516(data string) string {
	return r.Md532(data)[8:24]
}

// Md532 md5 32
func (r Rsae) Md532(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1 sha1
func (r Rsae) SHA1(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

// SHA256 sha256
func (r Rsae) SHA256(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

/**
*	解析公钥
*	@parmar:1、 publicKeyData []byte 需要解密公钥
*
*	@return:1、 *rsa.PublicKey
 */
func ParsePublicKey(publicKeyData []byte) *rsa.PublicKey {

	pubBlock, _ := pem.Decode(publicKeyData)

	pubKeyValue, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		fmt.Println("Parse PublicKey err:%s", err.Error())
		return nil
	}
	pubKey := pubKeyValue.(*rsa.PublicKey)

	return pubKey
}

/**
*	解析私钥
*	@parmar:1、 privateKeyData []byte 需要解密私钥
*
*	@return:1、 *rsa.PrivateKey
 */
func ParsePrivateKey(privateKeyData []byte) *rsa.PrivateKey {

	priBlock, _ := pem.Decode(privateKeyData)

	priKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
	if err != nil {
		fmt.Println("Parse PrivateKey err:%s", err.Error())
		return nil
	}

	return priKey
}

/**
*	PKCS1v15加密
*	注意：用这个函数加密纯文本是很危险的，尽量使用下面的EncryptOAEP方法
*	@parmar:1、 priKey *rsa.PublicKey 加密公钥
*			2、 msg []byte 要加密内容
*
*	@return:1、 []byte 加密后的数据
 */
func EncryptPKCS1v15(pubKey *rsa.PublicKey, msg []byte) []byte {

	encryptPKCS15, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
	if err != nil {
		fmt.Println("rsa EncryptPKCS1v15 err:%s", err.Error())
		return nil
	}

	return encryptPKCS15
}

/**
*	解密PKCS1v15加密的内容
*	@parmar:1、 priKey *rsa.PrivateKey 解密私钥
*			2、 ciphertext []byte 要解密内容
*
*	@return:1、 []byte 解密后的数据
 */
func DecryptPKCS1v15(priKey *rsa.PrivateKey, ciphertext []byte) []byte {

	decryptPKCS, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, ciphertext)
	if err != nil {
		fmt.Println("rsa DecryptPKCS1v15 err:%s", err.Error())
		return nil
	}

	return decryptPKCS
}

/**
*	EncryptOAEP加密
*	@parmar:1、 priKey *rsa.PublicKey 加密公钥
*			2、 msg []byte 要加密内容
*
*	@return:1、 []byte 加密后的数据
 */
func EncryptOAEP(pubKey *rsa.PublicKey, msg []byte) []byte {

	encryptOAEP, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubKey, msg, nil)
	if err != nil {
		fmt.Println("rsa EncryptOAEP err:%s", err.Error())
		return nil
	}

	return encryptOAEP
}

/**
*	解密RSA-OAEP方式加密后的内容
*	@parmar:1、 priKey *rsa.PrivateKey 解密私钥
*			2、 ciphertext []byte 要解密内容
*
*	@return:1、 []byte 解密后的数据
 */
func DecryptOAEP(priKey *rsa.PrivateKey, ciphertext []byte) []byte {

	decryptOAEP, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, priKey, ciphertext, nil)
	if err != nil {
		fmt.Println("rsa DecryptOAEP err:%s", err.Error())
		return nil
	}

	return decryptOAEP
}
