package types

import (
	"testing"

	"github.com/getlantern/testify/assert"
)

func TestEmailEquality(t *testing.T) {
	emails := []string{
		"abc@gmail.com",
		"AbC@GMaiL.COM",
		"a.bc@gmail.com",
		"a.b.c@gmail.com",
		"abc+123@gmail.com",
		"a.b.c+123@gmail.com",
		"abc@googlemail.com",
		"A.b.C+123@googlemail.com",
	}

	parsed := make([]string, 0, len(emails))
	for _, email := range emails {
		p, err := ParseEmail(email)
		assert.NoError(t, err, "Should be able to parse", email)
		assert.Equal(t, "abc@gmail.com", p.String(), "Address should normalize to abc@gmail.com")
	}

	for i := 0; i < len(parsed); i++ {
		a := parsed[i-1]
		b := parsed[i]
		assert.Equal(t, a, b, "All parsed emails should be equal")
	}
}
