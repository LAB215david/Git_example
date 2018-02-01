package main

import (
    "log"
    "fmt"
    "github.com/sideshow/apns2"
    "github.com/sideshow/apns2/certificate"
)

func main() {
    P12 := "PPhookDeviceVoIP.p12"
    P12_password := "mwnlmwnl"
    // IPhoneToken := "9366e08dfc9ce1229a332c916e91055f1eccb05dc3006088537a0bd0383ee4ce"
    // IPhoneToken := "c459fff0b2d75bb893034535ad63bffccfb5f0400963e3d43bf114435259e22c"
    // Certificate_Topic := "tw.com.linphpp.pphook.voip"

    // IPhoneToken := "e6bce2fa996c07c854f1960e613b7b0e4af57f1ea1ea9ebda7d987bb5cd28b29"
    // Certificate_Topic := "tw.com.ite.intercom.pphook.voip"

    IPhoneToken := "27b1adbb2ff4e374f2158a14f7559cfe57345030ab673175a80d0bc3a11fc93e"
    Certificate_Topic := "tw.com.devicenotification.pphook.voip"

    PushData := "TestData From Apple"
    FormatPushData := fmt.Sprintf("{\"aps\":{\"alert\":\"%s\"}}", PushData)
    
    cert, err := certificate.FromP12File(P12, P12_password)
    if err != nil {
    log.Fatal("Cert Error:", err)
    }

    notification := &apns2.Notification{}
    notification.DeviceToken = IPhoneToken
    notification.Topic = Certificate_Topic
    notification.Payload = []byte(FormatPushData)

    // client := apns2.NewClient(cert).Production()
    client := apns2.NewClient(cert).Development()

    res, err := client.Push(notification)

    if err != nil {
      log.Fatal("Error:", err)
    }

    fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}