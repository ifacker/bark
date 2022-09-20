package BarkNotification

import (
	"fmt"
	"testing"
)

func TestBark_SendMessage(t *testing.T) {
	bark := Bark{
		// 请求的URL，包括Key (必需要有）
		Url: "https://api.day.app/xxxxxxx",
		// 通知标题 (必需要有）
		Title: "title",
		// 通知内容
		Body: "body",
		// 推送铃声
		Sound: Sound_Alarm,
		// 是否保存通知消息
		IsArchive: Is_Archive_ON,
		// 通知图标
		Icon: "https://s1.ax1x.com/2022/09/16/vzIC9K.png",
		// 接受消息分组的组名
		GroupName: "test",
		// 时效性通知
		Level: Level_Active,
		// 通知 Url 跳转 如：https://www.baidu.com
		Jump2Url: "https://www.baidu.com",
		// 只复制 copy 参数到值
		Copy: "copyText",
		// 设置角标
		Badge: 1,
		// 自动复制 需要与 Copy 组合使用
		AutoCopy: AutoCopy_OFF,
	}
	a, err := bark.SendMessage() // 配置好 bark 之后直接发送消息即可
	fmt.Println(a, err)
}
