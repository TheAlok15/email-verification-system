package validators

import "github.com/TheAlok15/email-verification-system/internal/verifier/core"


type DisposableValidator struct{}
func (d DisposableValidator) Validate(ctx *core.VerificationContext) error {

	disposableDomain := map[string]bool{
		"tempmail.com" : true,
		"10mail.com" : true,
	}

	if disposableDomain[ctx.Domain] {
		ctx.Result.Disposable = true
	} else {
		ctx.Result.Disposable = false
	}

	return nil



}