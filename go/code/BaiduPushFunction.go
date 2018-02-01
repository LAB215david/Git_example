package main

import (
	"fmt"
	"github.com/ChangjunZhao/BaiduPushSDK-golang"
)

func main() {

	PushData := "TestData From Baidu"
	FormatPushData := fmt.Sprintf("{\"aps\": {  \"alert\":\"%s\",\"sound\":\"\",  \"badge\":0, \"content-available\":1}}", PushData)
	PushKey := "fum8a9CZlolyszaL60EECqLZ"
	SecretKey := "fd51RTR8telNqkvKIhw2eDscGuin7j0C"
	DeviceChannelID := "5285727420856693532"
	
	// 新建客户端
	client := push.NewClient(PushKey,SecretKey)
	
	// 构造请求
	request := &push.PushMsgToSingleDeviceRequest{ChannelId: DeviceChannelID, MsgType: 0, Message: FormatPushData, DeployStatus: 1}
	
	// 推送消息到指定客户端
	response, err := client.PushMsgToSingleDevice(*request)
	if err == nil {
		fmt.Println(response.MsgId)
	} else {
		fmt.Println(err)
	}
}