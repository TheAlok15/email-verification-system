package validators

import (
	"errors"
	"strings"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type DomainExtractor struct{}

func (v DomainExtractor) Validate(ctx *core.VerificationContext) error {

	parts := strings.Split(ctx.Email, "@")

	if len(parts) != 2 {
		return errors.New("invalid email format")
	}

	ctx.Domain = parts[1]

	return nil
}