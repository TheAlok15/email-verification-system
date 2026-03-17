package validators

import (
	"net"
	"strings"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type MXValidator struct{}

func (m MXValidator) Validate(ctx *core.VerificationContext) error {

	domain := strings.ToLower(ctx.Domain)
	mxRecord, err := net.LookupMX(domain)

	if err != nil || len(mxRecord) == 0 {
		ctx.Result.HasMX = false
		return nil
	} else {
		ctx.Result.HasMX = true
	}

	return nil

}
