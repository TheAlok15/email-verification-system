package validators

import (
	"net"
	"strings"

	"github.com/TheAlok15/email-verification-system/internal/verifier/core"
)

type DMARCValidator struct{}

func (d DMARCValidator) Validate(ctx *core.VerificationContext) error {

	domain := strings.ToLower(ctx.Domain)
	newDomain := "_dmarc." + domain

	records, err := net.LookupTXT(newDomain)

	if err != nil || len(records) == 0 {
		ctx.Result.HasDMARC = false
		return nil
	}


	for _, record := range records {
		if strings.Contains(strings.ToLower(record), "v=dmarc1") {
			ctx.Result.HasDMARC = true
			return nil
		}
	}
	ctx.Result.HasDMARC = false

	return nil
}
