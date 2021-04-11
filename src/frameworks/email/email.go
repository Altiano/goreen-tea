package email

type (
	Emailer interface {
		Send(email Email) error
	}

	Email struct {
		From string
		To   string
		Body string
	}
)
