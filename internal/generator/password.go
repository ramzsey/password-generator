package generator

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$+<?"

type Password struct {
	Short  string
	Medium string
	Long   string
}

func Generate() (*Password, error) {
	short, err := generatePassword(20)
	if err != nil {
		return nil, err
	}

	medium, err := generatePassword(30)
	if err != nil {
		return nil, err
	}

	long, err := generatePassword(40)
	if err != nil {
		return nil, err
	}

	return &Password{
		Short:  short,
		Medium: medium,
		Long:   long,
	}, nil
}

func generatePassword(length int) (string, error) {
	password := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}

	return string(password), nil
}

func GenerateSingle(length int) (string, error) {
	return generatePassword(length)
}
