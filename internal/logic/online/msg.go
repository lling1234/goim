package logic

// A和B在线单聊
type SendMsgReq struct {
	SendID           int64  `json:"sendid" description:"发送用户id"`
	ReceiveID        int64  `json:"recvid" description:"接收用户id"`
	SessionType      int64  `json:"sessionType" description:"会话类型 0-单聊 1-群聊"`
	SessionID        int64  `json:"sessionid" description:"会话id"`
	ContentType      int64  `json:"contentType" description:"消息类型"`
	MsgID            int64  `json:"msgid" description:"消息id"`
	MsgContent       []byte `json:"msgcontent" description:"消息内容"`
	GroupID          int64  `json:"groupid" description:"群id"`
	SenderPlatformID int64  `json:"senderPlatformID" description:"发送者平台id"`
	SendTime         int64  `json:"sendTime" description:"发送者平台id"`
	Ex               string `json:"ex" description:"拓展数据"`
}

type ReceiveMsgResp struct {
	SendID           int64  `json:"sendid" description:"发送用户id"`
	ReceiveID        int64  `json:"recvid" description:"接收用户id"`
	SessionType      int64  `json:"sessionType" description:"会话类型 0-单聊 1-群聊"`
	SessionID        int64  `json:"sessionid" description:"会话id"`
	ContentType      int64  `json:"contentType" description:"消息类型"`
	Seq              int64  `json:"seq" description:"消息序列"`
	MsgID            int64  `json:"msgid" description:"消息id"`
	MsgContent       []byte `json:"msgcontent" description:"消息内容"`
	GroupID          int64  `json:"groupid" description:"群id"`
	SenderPlatformID int64  `json:"senderPlatformID" description:"发送者平台id"`
	SendTime         int64  `json:"sendTime" description:"发送者平台id"`
	CreatTime        int64  `json:"createTime" description:"发送者平台id"`
	Ex string `json:"ex" description:"拓展数据"`
}
