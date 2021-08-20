package encrypt_test

import (
	"rentRoom/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestHashPassword(t *testing.T) {
	password := "secret"
	hash,error := encrypt.Hash(password)
	assert.Nil(t, error)
	assert.True(t, encrypt.ValidateHash(password, hash))
}
