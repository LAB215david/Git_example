package main

import (
	"fmt"
    "github.com/NaySoftware/go-fcm"
)

const (
	 serverKey = "AAAArCd01lE:APA91bFdPcvpUJnu7bL94zWH1XA2nvLWko4Jbl7a-MgcIiUyISDGI2Gu-SWVaynnnr3Rhk-upxZJHBS8ybvkWe77xQSl_OdHxjmiKfQKmQeRy2JvK7nNswt-TN5DjSanWyPmGJv3py2s"
	 //Just change "alarms" part to your topic!
     topic = "/topics/alarms"
)

func main() {

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(topic, data)


	status, err := c.Send()


	if err == nil {
    status.PrintResults()
	} else {
		fmt.Println(err)
	}

}