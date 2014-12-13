// lib Общие вспомогательные функци
// Работа со строками
package str

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

type String struct {
}

func NewString() *String {
	var self = new(String)
	return self
}

// CreatePassword make random password
func (self *String) CreatePassword() string {
	c := 10
	b := make([]byte, c)
	n, err := io.ReadFull(rand.Reader, b)
	if n != len(b) || err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%x", b)
}

// CreatePasswordHash make password hash
func (self *String) CreatePasswordHash(password string) string {
	shaCoo := sha256.New()
	shaCoo.Write([]byte(password))
	return fmt.Sprintf("%x", shaCoo.Sum(nil))
}

// CheckSum make check summ
func (self *String) CheckSum(data []byte) (checkSum string) {
	shaCoo := sha1.New()
	shaCoo.Write(data)
	return fmt.Sprintf("%x", shaCoo.Sum(nil))
}

func (self *String) Base64Encode(src []byte) (res string) {
	return base64.StdEncoding.EncodeToString(src)
}

func (self *String) Base64Decode(value string) (src []byte) {
	if data, err := base64.StdEncoding.DecodeString(value); err == nil {
		return data
	}
	return
}

func (self *String) Capitalize(str string) (ret string) {
	var tmp string
	tmp = strings.ToLower(str)
	firstRune, n := utf8.DecodeRuneInString(tmp)
	ret = string(unicode.ToUpper(firstRune)) + tmp[n:]
	return
}
