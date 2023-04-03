package party

import (
	"fmt"
	kavenegar "github.com/kavenegar/kavenegar-go"
	"github.com/pkg/errors"
	"shopChallenge/thirdparty/domain"
)

type KavenegarSMS struct {
	client *kavenegar.Kavenegar
	next   domain.SMS
}

func initKavenegarSMS(apiKey string) domain.SMS {
	c := kavenegar.New(apiKey)
	return &KavenegarSMS{client: c}
}

func (k *KavenegarSMS) SendMessage(message, receptor string) error {
	sender := ""
	receptors := []string{receptor}
	if res, err := k.client.Message.Send(
		sender, receptors, message, nil); err != nil {
		return errors.New(fmt.Sprint(
			"Kavenegar cannot send message error:%s",
			err.Error(),
		))
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
		}
	}
	if k.next != nil {
		return k.next.SendMessage(message, receptor)
	}
	return nil
}

func (k *KavenegarSMS) SetNext(next domain.SMS) {
	k.next = next
}

func (k *KavenegarSMS) GetNext() domain.SMS {
	return k.next
}
