package party

import (
	"fmt"
	dm "shopChallenge/domain"
	"shopChallenge/thirdparty/domain"
)

func InitSMS(smsConfig domain.Config) domain.SMS {
	kavenegar := initKavenegarSMS(smsConfig.Kavenegar.Apikey)
	ghasedak := initGhasedakSMS(smsConfig.Ghasedak.Apikey)
	sms := initiatorSMS{}
	if smsConfig.Kavenegar.Active {
		sms.SetNext(kavenegar)
		fmt.Println("kave created")
	}
	if smsConfig.Ghasedak.Active {
		fmt.Println("ghasedak created")
		sms.SetNext(ghasedak)
	}
	return &sms
}

type initiatorSMS struct {
	Next domain.SMS
}

func (i *initiatorSMS) SendMessage(message, receptor string) error {
	if i.Next != nil {
		return i.Next.SendMessage(message, receptor)
	}
	return dm.ErrNoActiveSMSService
}

func (i *initiatorSMS) SetNext(next domain.SMS) {
	for i.Next != nil {
		i.Next = i.Next.GetNext()
	}
	i.Next = next
}

func (i *initiatorSMS) GetNext() domain.SMS {
	return i.Next
}
