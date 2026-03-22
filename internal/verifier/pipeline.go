package verifier

import (
	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
	"github.com/TheAlok15/email-verification-system/internal/verifier/validators"
)

func VerifyEmail(email string) (*core.VerificationContext, error) {

	ctx := &core.VerificationContext{
		Email: email,
	}

	validatorsList := []core.Validator{
		validators.SyntaxValidator{},
		validators.DomainExtractor{},
		validators.DisposableValidator{},
		validators.RoleBasedValidator{},
		validators.MXValidator{},
		validators.SPFValidator{},
		validators.DMARCValidator{},
		validators.SMTPValidator{},
	}

	for _, v := range validatorsList {
		if err := v.Validate(ctx); err != nil {
			return nil, err
		}
	}

	return ctx, nil
}
