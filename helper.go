package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetHashMac(mackey *string, databody *string) *string {
	secret := *mackey
	data := *databody

	//fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	//fmt.Println("Result: " + sha)

	return &sha
}

func Getauthentication(secretkey *string, dataEndcode *string) *string {
	srcData := *dataEndcode
	key := []byte(*secretkey)
	//Test encryption
	encData, err := ECBEncrypt([]byte(srcData), (key))
	if err != nil {
		return nil
	}

	// //Test decryption
	// decData, err := ECBDecrypt(encData, key)
	// if err != nil {

	// 	return nil
	// }

	xdata := string(encData)
	return &xdata

}

func ECBDecrypt(crypted, key []byte) ([]byte, error) {
	if !validKey(key) {
		return nil, fmt.Errorf("the length of the secret key is wrong, the current incoming length is% d", len(key))
	}
	if len(crypted) < 1 {
		return nil, fmt.Errorf("source data length cannot be 0")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypted)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("the source data length must be an integer multiple of% D, the current length is% d", block.BlockSize(), len(crypted))
	}
	var dst []byte
	tmpData := make([]byte, block.BlockSize())

	for index := 0; index < len(crypted); index += block.BlockSize() {
		block.Decrypt(tmpData, crypted[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}

	dst, err = PKCS5UnPadding(dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func ECBEncrypt(src, key []byte) ([]byte, error) {
	if !validKey(key) {
		return nil, fmt.Errorf("the length of the secret key is wrong, the current incoming length is% d", len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) < 1 {
		return nil, fmt.Errorf("source data length cannot be 0")
	}
	src = PKCS5Padding(src, block.BlockSize())
	if len(src)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("the source data length must be an integer multiple of% D, the current length is% d", block.BlockSize(), len(src))
	}
	var dst []byte
	tmpData := make([]byte, block.BlockSize())
	for index := 0; index < len(src); index += block.BlockSize() {
		block.Encrypt(tmpData, src[index:index+block.BlockSize()])
		dst = append(dst, tmpData...)
	}
	return dst, nil
}

//Pkcs5 filling
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//Remove pkcs5 filling
func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])

	if length < unpadding {
		return nil, fmt.Errorf("invalid unpadding length")
	}
	return origData[:(length - unpadding)], nil
}

//Key length verification
func validKey(key []byte) bool {
	k := len(key)
	switch k {
	default:
		return false
	case 16, 24, 32:
		return true
	}
}

func EncodeString(data string) string {

	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	return encoded

}
func EncondeFile(filename string, errorLog *log.Logger) *string {

	// Open file on disk.
	//f, _ := os.Open("./file/01 Aset.txt")
	f, err := os.Open("./" + filename)

	if err != nil {
		errorLog.Fatal("Can not Open file "+filename, " --> ", err.Error())
	}

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		errorLog.Fatal("Can not Read file "+filename+" --> ", err.Error())
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return &encoded
}
