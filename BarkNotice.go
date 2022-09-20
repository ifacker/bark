package BarkNotification

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ifacker/myutil"
	"net/http"
)

// 发送消息
func (b *Bark) SendMessage() (string, error) {
	lastUrl := b.Url[len(b.Url)-1:]
	if lastUrl == "/" {
		urltmp := b.Url[:len(b.Url)-1]
		b.Url = urltmp
	}
	sendUrl := fmt.Sprintf("%s/%s/%s", b.Url, b.Title, b.Body)
	sendUrl += fmt.Sprintf("?%s", b.andandand())
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, sendUrl, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	bodyByte, err := myutil.AutoReaderBody2Byte(resp)
	if err != nil {
		return "", err
	}
	respjson := &respJson{}
	json.Unmarshal(bodyByte, respjson)
	if respjson.Code != 200 {
		return "", errors.New("发送的请求有误")
	}
	return string(bodyByte), nil
}

// 对读取对参数进行自动识别组装
func (b *Bark) andandand() string {
	result := ""
	switch {
	case b.Sound != "":
		result += fmt.Sprintf("sound=%s&", b.Sound)
	case b.IsArchive == Is_Archive_ON:
		result += fmt.Sprintf("isArchive=%d&", Is_Archive_ON)
	case b.IsArchive == Is_Archive_OFF:
		result += fmt.Sprintf("isArchive=%d&", Is_Archive_OFF)
	case b.Icon != "":
		result += fmt.Sprintf("icon=%s&", b.Icon)
	case b.GroupName != "":
		result += fmt.Sprintf("group=%s&", b.GroupName)
	case b.Level != "":
		result += fmt.Sprintf("level=%s&", b.Level)
	case b.Jump2Url != "":
		result += fmt.Sprintf("url=%s&", b.Jump2Url)
	case b.Copy != "":
		result += fmt.Sprintf("copy=%s&", b.Copy)
	case b.Badge != 0:
		result += fmt.Sprintf("icon=%d&", b.Badge)
	case b.AutoCopy == AutoCopy_ON && b.Copy != "":
		result += fmt.Sprintf("autoCopy=%d&", AutoCopy_ON)
	}
	return result[:len(result)-1]
}
