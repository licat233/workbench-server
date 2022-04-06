package cryptx

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const (
	SECRET = "licat2333333"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

/** MD5V
 * description: md5加密
 * param: str []byte, b ...byte
 * return: string
 */
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func ToMD5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
