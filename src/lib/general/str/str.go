// Библиотека для работы со строками
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

// Тип для агрегирования методов работы со строками
type String struct {
}

// Конструктор
//	- *String объект агрегатор функций-методов
func NewString() *String {
	var self = new(String)
	return self
}

// Make random password
//	- string сгенерированный пароль
func (self *String) CreatePassword() string {
	c := 10
	b := make([]byte, c)
	n, err := io.ReadFull(rand.Reader, b)
	if n != len(b) || err != nil {
		fmt.Println("error:", err)
	}
	return fmt.Sprintf("%x", b)
}

// Make password hash
//	+ password string строка хеш которой нужно получить (как правило для различных паролей)
//	- string хеш строка
func (self *String) CreatePasswordHash(password string) string {
	shaCoo := sha256.New()
	shaCoo.Write([]byte(password))
	return fmt.Sprintf("%x", shaCoo.Sum(nil))
}

// Make check summ
//	+ data []byte любые данные
//	- string контрольная сумма
func (self *String) CheckSum(data []byte) (checkSum string) {
	shaCoo := sha1.New()
	shaCoo.Write(data)
	return fmt.Sprintf("%x", shaCoo.Sum(nil))
}

// Кодировка в Base64
//	+ src []byte входные данные для кодирования
//	- string закодирвоанная строка
func (self *String) Base64Encode(src []byte) (res string) {
	return base64.StdEncoding.EncodeToString(src)
}

// Обратное кодирование из Base64
//	- value string закодирвоанная строка
//	+ []byte раскодированные данные
func (self *String) Base64Decode(value string) (src []byte) {
	if data, err := base64.StdEncoding.DecodeString(value); err == nil {
		return data
	}
	return
}

// Первый символ первого слова строки в верхний регистр
//	+ str string входная строка
//	- string полученная строка
func (self *String) Capitalize(str string) (ret string) {
	var tmp string
	tmp = strings.ToLower(str)
	firstRune, n := utf8.DecodeRuneInString(tmp)
	ret = string(unicode.ToUpper(firstRune)) + tmp[n:]
	return
}

////
