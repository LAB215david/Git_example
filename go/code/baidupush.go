package main

import (
	"fmt"
	"github.com/ChangjunZhao/BaiduPushSDK-golang"
	"time"
)

func main() {
	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

 

	//格式化为字符串,tm为Time类型

	tm := time.Unix(timestamp, 0)

	fmt.Println(tm.Format("2006/01/02, 15 04 05"))
	RightNowTime := tm.Format("2006/01/02, 15:04:05")
	PushData := fmt.Sprintf("{\"aps\": {  \"alert\":\"%s\",\"sound\":\"\",  \"badge\":0, \"content-available\":1}}", RightNowTime)

	// 新建客户端
	client := push.NewClient("fum8a9CZlolyszaL60EECqLZ","fd51RTR8telNqkvKIhw2eDscGuin7j0C")
	// client := push.NewClient("6BlG442WhkRax3zA9MpOwGln","EMvnwODYhN6yHAd54UpB6RhBL7CwzVfb")
	// 构造请求
	request := &push.PushMsgToSingleDeviceRequest{ChannelId: "5285727420856693532", MsgType: 0, Message: PushData, DeployStatus: 1}
	fmt.Println(request)
	// request := &push.PushMsgToSingleDeviceRequest{ChannelId: "4248305206004072057", MsgType: 0, Message: "TestBaiduSDK"}
	// 推送消息到指定客户端
	response, err := client.PushMsgToSingleDevice(*request)
	if err == nil {
		fmt.Println(response.MsgId)
	} else {
		fmt.Println(err)
	}
}