package generator

import (
	"crypto/rand"
	"encoding/binary"
)

var charset = func() string {
	chars := make([]byte, 94)
	for i := range chars {
		chars[i] = byte(i + 33)
	}
	return string(chars)
}()

const randomNumberRange = uint64(1) << 32

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
	charsetLen := len(charset)
	unbiasedLimit := uint32(randomNumberRange - (randomNumberRange % uint64(charsetLen)))
	password := make([]byte, 0, length)
	randomValues := make([]byte, 4*length)

	for len(password) < length {
		if _, err := rand.Read(randomValues); err != nil {
			return "", err
		}

		for i := 0; i < length && len(password) < length; i++ {
			value := binary.LittleEndian.Uint32(randomValues[i*4 : (i+1)*4])
			if value < unbiasedLimit {
				password = append(password, charset[value%uint32(charsetLen)])
			}
		}
	}

	return string(password), nil
}

func GenerateSingle(length int) (string, error) {
	return generatePassword(length)
}
