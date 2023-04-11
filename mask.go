package masking

import (
	"fmt"
	"strings"
)

const charMasking = "*"

type mask struct {
	str string
}

func (m *mask) overlay(start, end int) string {
	l := len([]rune(m.str))
	
	// when characters are empty
	if l < 1 {
		return ""
	}
	
	// when start masking is lower than 1
	if start < 1 {
		start = 1
	}
	
	// when start masking is more than len of char
	if start > l {
		start = l
	}
	
	// when end masking is lower than zero
	if end < 0 {
		end = 0
	}
	
	// when end masking is more than len of char
	if end > l {
		end = l
	}
	
	if start > end {
		tmp := start
		start = end
		end = tmp
	}
	
	return m.str[:start] + strings.Repeat(charMasking, end-start) + m.str[end:]
}

func (m *mask) defaultMasking() string {
	start := 3
	l := len(m.str)
	end := l - 2
	
	if l < 4 {
		end = l
	}
	
	if l < 6 {
		start = 1
	}
	
	return m.overlay(start, end)
}

func (m *mask) password() string {
	return strings.Repeat(charMasking, 10)
}

func (m *mask) creditCard() string {
	return m.overlay(2, len(m.str)-4)
}

func (m *mask) email() string {
	s := strings.Split(m.str, "@")
	if len(s) < 2 || len(s) > 2 {
		return ""
	}
	
	m.str = s[0]
	e := m.overlay(2, len(m.str)-1)
	m.str = s[1]
	d := m.overlay(2, len(m.str)-2)
	
	return fmt.Sprintf("%s@%s", e, d)
}

func (m *mask) phoneNumber() string {
	return m.overlay(5, len(m.str)-2)
}

func String(maskType Type, str string) (val string) {
	m := &mask{str}
	
	switch maskType {
	case Password:
		val = m.password()
	case CreditCard:
		val = m.creditCard()
	case Email:
		val = m.email()
	case PhoneNumber:
		val = m.phoneNumber()
	default:
		val = m.defaultMasking()
	}
	
	return
}
