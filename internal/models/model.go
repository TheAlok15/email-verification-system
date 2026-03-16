package models

import "time"

type JobStatus string
const (
	JobPending    JobStatus = "pending"
	JobProcessing JobStatus = "processing"
	JobCompleted  JobStatus = "completed"
	JobFailed     JobStatus = "failed"
)

type Classification string
const (
	ClassValid      Classification = "valid"
	ClassInvalid    Classification = "invalid"
	ClassCatchAll   Classification = "catch_all"
	ClassDisposable Classification = "disposable"
	ClassUnknown    Classification = "unknown"
)

type DomainInfo struct {
	Domain string

	HasMX bool
	HasSPF bool
	HasDMARC bool

	CatchAll bool

	CheckedAt time.Time
}

type Job struct {
	ID        string
	Email     string
	Status    JobStatus
	Attempts  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Result struct {
	JobID  string
	Domain string

	HasMX    bool
	HasSPF   bool
	HasDMARC bool

	SMTPValid   bool
	SMTPCode    int
	SMTPMessage string

	Disposable bool
	CatchAll   bool
	RoleBased  bool

	Score          int
	Classification Classification
	CheckedAt time.Time
}

type Input struct {
	Email string `json:"email"`
}