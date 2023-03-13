package util

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

const FromPhoneNumber = "+12763228670"

func NotifyUser(phoneNumber string, message string) error {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom(FromPhoneNumber)
	params.SetTo("+1" + phoneNumber)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("twilio error: " + err.Error())
		return err
	}

	return nil

}
