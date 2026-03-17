package validators

import (
	"strings"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type RoleBasedValidator struct{}

func (r RoleBasedValidator) Validate(ctx *core.VerificationContext) error {

	roleBasedList := map[string]bool{
		"admin":   true,
		"support": true,
		"info":    true,
		"sales":   true,
		"contact": true,
	}

	parts := strings.Split(ctx.Email, "@")
	if len(parts) != 2 {
		return nil
	}

	local := strings.ToLower(parts[0])

	if roleBasedList[local] {
		ctx.Result.RoleBased = true
	} else {
		ctx.Result.RoleBased = false
	}

	return nil

}
