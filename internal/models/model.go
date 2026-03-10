package models

import "time"

type Job struct {
	ID        string
	Email     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Result struct {
	JobID string

	HasMX    bool
	HasSPF   bool
	HasDMARC bool

	SMTPValid bool
	Disposable bool
	CatchAll bool
	RoleBased bool

	Score int
	Classification string

	CheckedAt time.Time
}

type Input struct{
	Email string `json:"email"`
}