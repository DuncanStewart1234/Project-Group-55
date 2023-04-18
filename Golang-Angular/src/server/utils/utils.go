// This package contains a utility function for the sqlite databases
package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"html"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GetDB returns the specified database for use in other packages
func GetDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database " + path)
	}
	return db
}

func CheckIfEmptyOrTooLong(msg string) (error) {
	if msg == "" || len(msg) > 256 {
		return errors.New("string cannot be empty nor longer than 256 chars")
	}

	return nil
}

func Fix_Username(uname string) (string) {
	return html.EscapeString(strings.TrimSpace(uname))
}

func HashPasswrd(password string) []byte {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to hash password")
	}

	return hashedPass
}

func CheckHashedPasswrd(password string, hash []byte) (error) {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func CheckIfClassExists(cid string) (bool) {
	var result bool
	db := GetDB("src/server/databases/classes.db")
	err := db.Where("Class_ID = ?", cid).Find(&result).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func GetPrivKey() []byte {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	return ExportPrivateKey(key)
}

func ExportPrivateKey(privkey *rsa.PrivateKey) []byte {
    privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
    privkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PRIVATE KEY",
                    Bytes: privkey_bytes,
            },
    )
    return privkey_pem
}

func ImportPrivateKey(pep_priv string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(pep_priv))
    if block == nil {
            return nil, errors.New("cannot parse")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

func GeneratePrivKey() {
	if err := os.WriteFile("src/server/key.rsa", GetPrivKey(), 0700); err != nil {
        panic(err)
    }
}