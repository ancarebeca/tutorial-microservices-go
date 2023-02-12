package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordMatches(t *testing.T) {

	u := &User{
		Password: "$2a$12$WxHmNUM04pbozNGUJIFazefXFImJK/cQKmh.V1E2I5oMjWCx2HQLa",
	}

	plainText := "verysecret!"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), 12)
	fmt.Println(hashedPassword)
	ok, err := u.PasswordMatches(plainText)

	assert.True(t, ok)
	assert.NoError(t, err)
}
