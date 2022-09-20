## Bark 的通知工具 🎉🎉🎉
<img src="https://camo.githubusercontent.com/8415d8c1f8c5e1a34cdc8e41c56ecf7c66371433d2d9aa9b4c792108a60e0b8a/68747470733a2f2f7778332e73696e61696d672e636e2f6d773639302f303036306c6d37546c793167306e666e6a6a7862626a333073673073673735372e6a7067">  

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