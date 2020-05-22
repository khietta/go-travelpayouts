package travelr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data = map[string]interface{}{
		"r": "rdffsa",
		"e": map[string]interface{}{
			"6": "2222",
			"1": "rtrt",
		},
	}
)

func TestMakeSignatureString(t *testing.T) {
	result := ":rtrt:2222:rdffsa"
	rs := makeSignatureString(data)

	assert.NotEqual(t, rs, "")
	assert.EqualValues(t, result, rs)
}

func TestMakeSignatureToken(t *testing.T) {
	rs, err := MakeSignatureToken(token, data)

	assert.Nil(t, err)
	assert.NotEqual(t, rs, "")
}
