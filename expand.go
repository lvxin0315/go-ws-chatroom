package GoWsChatroom

//保存房间信息集合
var roomInfoList = make(map[string] *RoomInfo)

//房间拓展的内容
type RoomInfoExpand struct {}

//客户端拓展的内容
type ClientExpand struct {}

type Expand interface {
	//设置客户端拓展信息的回调
	ClientExpandFunc (client *Client)
	//设置房间拓展信息的回调
	RoomInfoExpandFunc (roomInfo *RoomInfo)
	//广播信息留存回调
	SendMessageDataCallbackFunc(md *MessageData)
}
