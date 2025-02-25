package authentication

import (
	"time"

	"github.com/google/uuid"
)

type rbac int

const (
	BASIC rbac = iota
	USER
	ANALYST
	ADMIN
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Sid      uuid.UUID `json:"sid"`
	Rbac     rbac      `json:"rbac"`
	Token    string    `json:"token"`
	IssuedAt time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
