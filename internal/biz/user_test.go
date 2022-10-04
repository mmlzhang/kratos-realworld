package biz

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	var a int = 10
	pwd := "sssssssadfa"
	s := hashPassword(pwd)
	spew.Dump(s)
	fmt.Println("-->> ", s, a)
	// assert.Equal(t, s, h)

	at := assert.New(t)
	at.True(true)

	err := bcrypt.CompareHashAndPassword([]byte(s), []byte(pwd))
	if err != nil {
		fmt.Println(err)
	}
}
