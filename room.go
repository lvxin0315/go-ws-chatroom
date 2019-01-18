package GoWsChatroom

import "errors"

type RoomInfo struct {
	RoomId string
	RoomInfoExpand
}


//根据roomId获取房间信息
func (r *RoomInfo) GetRoomInfoByRoomId(roomId string) error {
	if roomInfoList[roomId] == nil {
		return errors.New("invalid roomId")
	}
	r.RoomId = roomInfoList[roomId].RoomId
	r.RoomInfoExpand = roomInfoList[roomId].RoomInfoExpand
	return nil
}

//创建一个房间
func (r *RoomInfo) CreateRoom() error {
	if r.RoomId == "" {
		return errors.New("roomId is error")
	}
	r.GetRoomInfoByRoomId(r.RoomId)

	if roomInfoList[r.RoomId] != nil {
		return errors.New("roomId already pure")
	}
	roomInfoList[r.RoomId] = r
	return nil
}

//关闭一个房间
func (r *RoomInfo) CloseRoom() {
	roomInfoList[r.RoomId] = nil
	delete(roomInfoList,r.RoomId)
}
