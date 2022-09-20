## Bark 的通知工具 🎉🎉🎉
<img src="https://wx3.sinaimg.cn/mw690/0060lm7Tly1g0nfnjjxbbj30sg0sg757.jpg" width=200px height=200px />  

### 🍺用法：👇
#### SendMessage
```go
	bark := Bark{
		Url:       "https://api.day.app/xxxxxxxx",
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
	a, err := bark.SendMessage()  // 配置好 bark 之后直接发送消息即可
	fmt.Println(a, err)
	
	// 打印：
	// {"code":200,"message":"success","timestamp":1663581572} <nil>
```

### ⚠️注意：👇  
创建 Bark 时，至少需要 Url 和 Title 这两个字段，其他字段可以根据自己的需求进行添加