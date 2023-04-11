package masking

type Type int

const (
	// Password masking
	Password Type = iota
	// CreditCard masking
	CreditCard
	// Email masking
	Email
	// PhoneNumber masking
	PhoneNumber
)
