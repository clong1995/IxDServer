package util

import (
	. "IxDServer/common"
	. "IxDServer/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

//解密token
func UnMarshalToken(tokenKey string) (string, error) {
	// Byte array of the string
	tokenKey = string([]byte(tokenKey)[len(CONF.AuthTokenStart):])
	ciphertext, err := hex.DecodeString(tokenKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Key
	key := []byte(CONF.EncryptedKey)
	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		log.Printf("%s is too short", CONF.EncryptedKey)
		return "", err
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)
	unMarshal := string(ciphertext)
	arr := strings.Split(unMarshal, ".")
	//是否有足够的信息
	if len(arr) != 2 {
		log.Println(AUTH_STR)
		return "", fmt.Errorf(AUTH_STR)
	}
	//校验时间
	prevTime, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if time.Now().Unix()-prevTime > CONF.AuthTokenDeadline {
		return "", fmt.Errorf(AUTH_STR)
	}
	return arr[0], nil
}

//加密token
func MarshalToken(tokenKey string) (string, error) {
	// Byte array of the string
	plaintext := []byte(tokenKey + "." + strconv.FormatInt(time.Now().Unix(), 10))
	// Key
	key := []byte(CONF.EncryptedKey)
	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]
	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println(err)
		return "", err
	}
	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)
	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	//log.Println(fmt.Sprintf("%x", ciphertext))
	return CONF.AuthTokenStart + fmt.Sprintf("%x", ciphertext), nil
}
