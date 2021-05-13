package mocks

import (
	"github.com/golang/mock/gomock"
	emailMocks "gitlab.com/altiano/goreen-tea/src/frameworks/email/mocks"
)

type (
	Email struct {
		Manager *emailMocks.MockManager
	}
)

func setupEmail(ctrl *gomock.Controller) Email {
	return Email{
		Manager: emailMocks.NewMockManager(ctrl),
	}
}
