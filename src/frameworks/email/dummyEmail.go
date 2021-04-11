package email

import (
	"gitlab.com/altiano/golang-boilerplate/src/shared"
)

type (
	dummy struct {
		config shared.Config
	}
)

func NewManager(config shared.Config) Emailer {
	return &dummy{
		config: config,
	}
}

func (m *dummy) Send(email Email) error {
	return nil
}
