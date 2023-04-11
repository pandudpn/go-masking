package masking_test

import (
	"github.com/pandudpn/go-masking"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestString_MaskPassword(t *testing.T) {
	var pwd = "123456"
	
	result := masking.String(masking.Password, pwd)
	expected := strings.Repeat("*", 10)
	
	assert.Equal(t, expected, result)
}

func TestString_MaskCreditCard(t *testing.T) {
	var pwd = "4000000000000002"
	
	result := masking.String(masking.CreditCard, pwd)
	expected := "40**********0002"
	
	assert.Equal(t, expected, result)
}

func TestString_MaskPhoneNumber(t *testing.T) {
	var pwd = "6283875181609"
	
	result := masking.String(masking.PhoneNumber, pwd)
	expected := "62838******09"
	
	assert.Equal(t, expected, result)
}

func TestString_MaskEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "success masking",
			email:    "pandu@pandudpn.id",
			expected: "pa**u@pa*******id",
		},
		{
			name:     "empty masking",
			email:    "pandudpn",
			expected: "",
		},
	}
	
	for _, tc := range tests {
		t.Run(
			tc.name, func(t *testing.T) {
				result := masking.String(masking.Email, tc.email)
				
				assert.Equal(t, tc.expected, result)
			},
		)
	}
}

func TestString_MaskDefault(t *testing.T) {
	tests := []struct {
		name     string
		val      string
		expected string
	}{
		{
			name:     "success masking",
			val:      "pandu dwi putra",
			expected: "pan**********ra",
		},
		{
			name:     "under 3 char",
			val:      "pan",
			expected: "p**",
		},
		{
			name:     "under 6 char",
			val:      "abcde",
			expected: "a**de",
		},
	}
	
	for _, tc := range tests {
		t.Run(
			tc.name, func(t *testing.T) {
				result := masking.String(10, tc.val)
				
				assert.Equal(t, tc.expected, result)
			},
		)
	}
}
