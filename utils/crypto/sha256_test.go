package crypto

import (
	"testing"

	"github.com/perfectbuii/gen-go-template/utils/helper"
	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	str := "something11111"
	hashed := HashSHA256(str)
	t.Log(hashed)
	isHash := helper.IsSHA256(hashed)
	assert.True(t, isHash)
}
