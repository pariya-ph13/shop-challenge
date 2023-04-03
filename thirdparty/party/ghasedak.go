package party

import (
	"fmt"
	ghasedak "github.com/ghasedakapi/ghasedak-go"
	"github.com/pkg/errors"
	"shopChallenge/thirdparty/domain"
)

type GhasedakSMS struct {
	ghasedak.Client
	next domain.SMS
}

func initGhasedakSMS(apiKey string) domain.SMS {
	c := ghasedak.NewClient(apiKey, "")
	return &GhasedakSMS{Client: c}
}

func (g *GhasedakSMS) SendMessage(message, receptor string) error {
	r := g.Client.Send(message, receptor)
	if !r.Success {
		return errors.New(fmt.Sprintf(
			"ghasedak cannot send message stat:%d",
			r.Code))
	}
	if g.next != nil {
		return g.next.SendMessage(message, receptor)
	}
	return nil
}

func (g *GhasedakSMS) SetNext(next domain.SMS) {
	g.next = next
}

func (g *GhasedakSMS) GetNext() domain.SMS {
	return g.next
}
