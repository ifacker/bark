package BarkNotification

// 通知铃声
const (
	Sound_Alarm              = "alarm"
	Sound_Anticipate         = "anticipate"
	Sound_Bell               = "bell"
	Sound_Birdsong           = "birdsong"
	Sound_Bloom              = "bloom"
	Sound_Calypso            = "calypso"
	Sound_Chime              = "chime"
	Sound_Choo               = "choo"
	Sound_Descent            = "descent"
	Sound_Electronic         = "electronic"
	Sound_Fanfare            = "fanfare"
	Sound_Glass              = "glass"
	Sound_Gotosleep          = "gotosleep"
	Sound_Healthnotification = "healthnotification"
	Sound_Lorn               = "horn"
	Sound_Ladder             = "ladder"
	Sound_Mailsent           = "mailsent"
	Sound_Minuet             = "minuet"
	Sound_Multiwayinvitation = "multiwayinvitation"
	Sound_Newmail            = "newmail"
	Sound_Newsflash          = "newsflash"
	Sound_Noir               = "noir"
	Sound_Paymentsuccess     = "paymentsuccess"
	Sound_Shake              = "shake"
	Sound_Sherwoodforest     = "sherwoodforest"
	Sound_Silence            = "silence"
	Sound_Spell              = "spell"
	Sound_Suspense           = "suspense"
	Sound_Telegraph          = "telegraph"
	Sound_Tiptoes            = "tiptoes"
	Sound_Typewriters        = "typewriters"
	Sound_Update             = "update"
)

// 保存通知消息
const (
	Is_Archive_ON  = 1 // 开启通知消息保存
	Is_Archive_OFF = 0 // 关闭通知消息保存
)

// 时效性通知
const (
	Level_Active        = "active"        // 默认值，系统会立即亮屏显示通知
	Level_TimeSensitive = "timeSensitive" // 时效性通知，可在专注状态下显示通知
	Level_Passive       = "passive"       // 仅将通知添加到通知列表，不会亮屏提醒
)

// 是否需要自动复制
const (
	AutoCopy_ON  = 1
	AutoCopy_OFF = 0
)

// Bark 数据类型
type Bark struct {
	// 请求的URL，包括Key (必需要有）
	Url string

	// 通知标题 (必需要有）
	Title string

	// 通知内容
	Body string

	// 推送铃声
	Sound string

	// 是否保存通知消息
	IsArchive int

	// 通知图标
	Icon string

	// 接受消息分组的组名
	GroupName string

	// 时效性通知
	Level string

	// 通知 Url 跳转 如：https://www.baidu.com
	Jump2Url string

	// 只复制 copy 参数到值
	Copy string

	// 设置角标
	Badge int

	// 自动复制 需要与 Copy 组合使用
	AutoCopy int
}

type respJson struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
