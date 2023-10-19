package encryption

import (
	"fmt"
	"strings"
)

type Encrypt struct {
}

var (
	ENCODING_PREFIX               = "ENC:"
	DETERMINISTIC_ENCODING_PREFIX = "DETERMINIST_ENC:"
)

func (encrypt *Encrypt) InitEncryption() error {
	var err error = nil
	return err
}

func (encrypt *Encrypt) CloseEncryption() error {
	var err error = nil
	return err
}

func (encrypt *Encrypt) Encrypt(clearText string, maxClearTextLen int) (string, error) {
	var err error = nil
	encodedText := ENCODING_PREFIX + clearText
	if len(encodedText) > maxClearTextLen {
		err = fmt.Errorf("%d %s", -5, "Buffer too small")
	}
	return encodedText, err
}

func (encrypt *Encrypt) EncryptDeterministic(clearText string, maxClearTextLen int) (string, error) {
	var err error = nil
	encodedText := DETERMINISTIC_ENCODING_PREFIX + clearText
	if len(encodedText) > maxClearTextLen {
		err = fmt.Errorf("%d %s", -5, "Buffer too small")
	}
	return encodedText, err
}

func (encrypt *Encrypt) Decrypt(cipherText string, maxClearTextLen int) (string, error) {
	var err error = nil
	if !strings.HasPrefix(cipherText, ENCODING_PREFIX) {
		return "", fmt.Errorf("%d %s", -1, "Invalid encrypted value")
	}
	clearText := strings.TrimPrefix(cipherText, ENCODING_PREFIX)
	if len(clearText) > maxClearTextLen {
		return clearText, fmt.Errorf("%d %sd", -5, "Buffer too small")
	}
	return clearText, err
}

func (encrypt *Encrypt) DecryptDeterministic(cipherText string, maxClearTextLen int) (string, error) {
	var err error = nil
	if !strings.HasPrefix(cipherText, ENCODING_PREFIX) {
		return "", fmt.Errorf("%d %s", -1, "Invalid encrypted value")
	}
	clearText := strings.TrimPrefix(cipherText, ENCODING_PREFIX)
	if len(clearText) > maxClearTextLen {
		return clearText, fmt.Errorf("%d %sd", -5, "Buffer too small")
	}
	return clearText, err
}
