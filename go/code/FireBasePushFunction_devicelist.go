package main

import (
	"fmt"
    "github.com/NaySoftware/go-fcm"
)

const (
	 serverKey = "AAAArCd01lE:APA91bFdPcvpUJnu7bL94zWH1XA2nvLWko4Jbl7a-MgcIiUyISDGI2Gu-SWVaynnnr3Rhk-upxZJHBS8ybvkWe77xQSl_OdHxjmiKfQKmQeRy2JvK7nNswt-TN5DjSanWyPmGJv3py2s"
)

func main() {

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

  ids := []string{
      "fPIZGMUhtNw:APA91bGutPwmAcVVq2ag749BFtfY6LZP-YNjfMfUlBUdKjkSp0NjZdx4vJQYzr0l1ufdZT2vq5yvU6QJOvfgtn9vjbFHKpsMSpbuKeg9P8bkd6cB13wEulKbAExCpQBaKIJ-odv41WCm",
  }


  // xds := []string{
  //     "token5",
  //     "token6",
  //     "token7",
  // }

  xds := []string{
      "dc5uQ1lS4t8:APA91bFIIEwdeni17ceGkBho33Eay9c4CzvIg1V9LwLP2X4H5oplG6N4aRjghnGS9q_JljU_M9L5qZUMHxkoeQstz7uTngkBKEBhalZJtTqJFo4wv2GiDvsuCMLh1BT34rb-zbMq9e-C",
  }

	c := fcm.NewFcmClient(serverKey)
    c.NewFcmRegIdsMsg(ids, data)
    c.AppendDevices(xds)

	status, err := c.Send()


	if err == nil {
    status.PrintResults()
	} else {
		fmt.Println(err)
	}

}