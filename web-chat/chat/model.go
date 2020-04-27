package chat

import "time"

// 用户在聊天室中的唯一id
type uid = string

const (
	EventTypeMsg    = "event-msg"    // 用户发言
	EventTypeSystem = "event-system" // 系统消息推送 如房间人数
	EventTypeJoin   = "event-join"   // 用户加入
	EventTypeTyping = "event-typing" // 用户正在输入
	EventTypeLeave  = "event-leave"  // 用户退出
	EventTypeImage  = "event-image"  // todo 消息图片
)

// 聊天室事件定义
type Event struct {
	Type      string `json:"type"`      // 事件类型
	User      string `json:"user"`      // 用户名
	Timestamp int64  `json:"timestamp"` // 时间戳
	Text      string `json:"text"`      // 事件内容
	UserCount int    `json:"userCount"` // 房间用户数
}

// 聊天室事件初始化
func NewEvent(typ string, user, msg string) Event {
	return Event{
		Type:      typ,
		User:      user,
		Timestamp: time.Now().UnixNano(),
		Text:      msg,
	}
}

// 用户订阅
type Subscription struct {
	Id       string       // 用户在聊天室中的id
	Username string       // 用户名
	Pipe     <-chan Event // 事件接收通道 用户从该通道接收消息
	EmitChn  chan Event   // 用户消息推送通道
	LeaveChn chan uid     // 用户离开时间推送
}

func (s *Subscription) Leave() {
	s.LeaveChn <- s.Id // 将用户从聊天室列表中移除
}

func (s *Subscription) Say(message string) {
	s.EmitChn <- NewEvent(EventTypeMsg, s.Username, message)
}
