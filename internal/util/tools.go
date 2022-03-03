package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

//对字符串进行MD5加密
func EncodeMD5(inputStr string) string {
	h := md5.New()
	h.Write([]byte(inputStr)) // 需要加密的字符串为 sharejs.com
	str := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return str
}

//base64字符加密
func Base64EncodeWithString(input string) string {

	return base64.StdEncoding.EncodeToString([]byte(input))
}

//base64字符解密
func Base64DecodeWithString(input string) (string, error) {

	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

// 验证密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
