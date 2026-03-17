package core

import "github.com/TheAlok15/email-verification-system/internal/models"


type VerificationContext struct{
	Email string
	Domain string
	Result models.Result
}

type Validator interface{
	Validate(ctx *VerificationContext) error
}


