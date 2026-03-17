package validators

import (
	"errors"
	"net/mail"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type SyntaxValidator struct{}

func (v SyntaxValidator) Validate(ctx *core.VerificationContext) error {

	_, err := mail.ParseAddress(ctx.Email)
	if err != nil {
		return errors.New("invalid email syntax")
	}

	return nil
}