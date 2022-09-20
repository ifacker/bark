package BarkNotification

import (
	"fmt"
	"testing"
)

func TestBark_SendMessage(t *testing.T) {
	bark := Bark{
		Url:       "https://api.day.app/xxxxxxx",
		Title:     "title",
		Body:      "body",
		Sound:     Sound_Alarm,
		IsArchive: Is_Archive_ON,
		Icon:      "https://s1.ax1x.com/2022/09/16/vzIC9K.png",
		GroupName: "test",
		Level:     Level_Active,
		Jump2Url:  "https://www.baidu.com",
		Copy:      "copyText",
		Badge:     1,
		AutoCopy:  AutoCopy_OFF,
	}
	a, err := bark.SendMessage()
	fmt.Println(a, err)
}
