package aurelia

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	uuid "github.com/satori/go.uuid"
)

type Hasher struct {
	hash       string
	credential string
	key        string
}

func Authenticate(credential, key, hash string) bool {
	hasher := Hasher{
		hash:       hash,
		credential: credential,
		key:        key,
	}

	return hasher.Authenticate()
}

func Hash(credential, key string) string {
	hasher := Hasher{
		credential: credential,
		key:        key,
	}

	return hasher.Hash()
}

func (self Hasher) Hash() string {

	u, _ := uuid.NewV4()

	randomChar := u.String()
	randomChar2 := u.String()

	hmacHash := hmac.New(sha1.New, []byte(self.key+randomChar+randomChar2))
	hmacHash.Write([]byte(randomChar + randomChar2))

	salt := hmacHash.Sum(nil)

	hmacHash2 := hmac.New(sha1.New, []byte(self.credential+self.key))
	hmacHash2.Write([]byte(fmt.Sprintf("%x", salt)))

	userUnique := hmacHash2.Sum(nil)

	dk := pbkdf2.Key(userUnique, []byte(fmt.Sprintf("%x", salt)), 4096, 32, sha1.New)

	return fmt.Sprintf("AURELIA_%x.%x.UC_%x", salt, userUnique, dk)
}

func (self Hasher) Authenticate() bool {
	salt := self.getSalt()
	signature := self.getSignature()

	hmacHash2 := hmac.New(sha1.New, []byte(self.credential+self.key))
	hmacHash2.Write([]byte(salt))

	userUnique := hmacHash2.Sum(nil)

	dk := pbkdf2.Key(userUnique, []byte(salt), 4096, 32, sha1.New)

	if fmt.Sprintf("%x", dk) == signature {
		return true
	}

	return false
}

func (self Hasher) getSalt() string {
	toBeSalt := strings.Split(self.hash, ".")[0]
	salt := strings.Split(toBeSalt, "_")

	return salt[1]
}

func (self Hasher) getSignature() string {
	return strings.Split(strings.Split(self.hash, ".")[2], "_")[1]
}
