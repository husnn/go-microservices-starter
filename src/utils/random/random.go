package random

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

func Int63() (int64, error) {
	max := big.NewInt(1)
	max.Lsh(max, 63)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

const digitChars = "1234567890"

func Digits(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(digitChars)
	for i := 0; i < length; i++ {
		buffer[i] = digitChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func Token(size int) (string, error) {
	key := make([]byte, size)
	n, err := io.ReadFull(rand.Reader, key)
	if n != size || err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
