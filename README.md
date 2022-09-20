## Bark 的通知工具 🎉🎉🎉
<img src="https://wx3.sinaimg.cn/mw690/0060lm7Tly1g0nfnjjxbbj30sg0sg757.jpg" width=200px height=200px />  

### 🍺用法：👇
#### SendMessage
```go
	bark := Bark{
		Url:       "https://api.day.app/xxxxxxxx", // 请求的URL，包括Key (必需要有）
		Title:     "title", // 通知标题 (必需要有）
		Body:      "body", // 通知内容
		Sound:     Sound_Alarm, // 推送铃声
		IsArchive: Is_Archive_ON, // 是否保存通知消息
		Icon:      "https://s1.ax1x.com/2022/09/16/vzIC9K.png", // 通知图标
		GroupName: "test", // 接受消息分组的组名
		Level:     Level_Active, // 时效性通知
		Jump2Url:  "https://www.baidu.com", // 通知 Url 跳转 如：https://www.baidu.com 
		Copy:      "copyText", // 只复制 copy 参数到值
		Badge:     1, // 设置角标
		AutoCopy:  AutoCopy_OFF, // 自动复制 需要与 Copy 组合使用
	}
	a, err := bark.SendMessage()  // 配置好 bark 之后直接发送消息即可
	fmt.Println(a, err)
	
	// 打印：
	// {"code":200,"message":"success","timestamp":1663581572} <nil>
```

### ⚠️注意：👇  
创建 Bark 时，至少需要 Url 和 Title 这两个字段，其他字段可以根据自己的需求进行添加