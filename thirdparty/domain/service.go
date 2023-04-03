package domain

type SMS interface {
	SendMessage(message, receptor string) error
	SetNext(next SMS)
	GetNext() SMS
}
