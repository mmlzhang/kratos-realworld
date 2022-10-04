package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJwtToken(t *testing.T) {
	token, err := GenerateJwtToken()
	if err != nil {
		panic(err)
	}
	spew.Dump(token)
	assert.True(t, true)

}
