package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

// File file
type File struct{}

// NewFile new file
func NewFile() *File {
	return &File{}
}

// Read file read
// date 2016-12-31
// author andy.jiang
func (f File) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// Write file write
func (f File) Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

/**
*	genRsaKey file
*	@parmar:
*	@return:
 */
func RsaFile() {

	var bits int
	bits = 2048
	if err := GenRsaKey(bits); err != nil {
		log.Fatal("密钥文件生成失败！")
	}
	log.Println("密钥文件生成成功！")

}

/**
*	read RsaKey
*	@parmar:
*	@return:1、 []byte, []byte :publicKey,privateKey
 */
func ReadKey() ([]byte, []byte) {

	publicKey, err := ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}

	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
	return publicKey, privateKey
}

/**
*	生成RsaKey
*	@parmar:1、 bits int
*	@return:1、 error
 */
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "private",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "public",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
