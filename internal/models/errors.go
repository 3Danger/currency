package models

type Error struct {
	Message string
}

func (e *Error) Error() string { return e.Message }

var (
	ErrCurrencyNotFound               = &Error{Message: "currency not found"}
	ErrFiatToFiatConvertForbidden     = &Error{Message: "fiat to fiat convert is forbidden"}
	ErrCryptoToCryptoConvertForbidden = &Error{Message: "crypto to crypto convert is forbidden"}
	ErrCodeInvalid                    = &Error{Message: "unknown code"}
	ErrCurrencyIsDeprecated           = &Error{Message: "currency is deprecated"}
)
