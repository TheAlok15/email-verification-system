package validators

import (
	"net"
	"strings"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type SPFValidator struct{}

func (s SPFValidator) Validate(ctx *core.VerificationContext) error {

	domain := strings.ToLower(ctx.Domain)

	records, err := net.LookupTXT(domain)

	if err != nil || len(records) == 0 {
		ctx.Result.HasSPF = false
		return nil
	}
	// because contains need string not slice of string that why we go fo range
	// check := strings.Contains(records, "v=spf1")

	for _, record := range records {
		if strings.Contains(strings.ToLower(record), "v=spf1") {
			ctx.Result.HasSPF = true
			return nil
		}
	}
	ctx.Result.HasSPF = false

	return nil
}
