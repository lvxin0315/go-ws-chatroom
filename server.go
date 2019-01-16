package GoWsChatroom

import (
	"net/http"
	"log"
	"errors"
)

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request,roomInfoExpandFunc RoomInfoExpandFunc,clientExpandFunc ClientExpandFunc) (error) {
	r.ParseForm()
	roomId := r.Form["room_id"][0]
	if roomId == ""{
		log.Println("get room_id error")
		return errors.New("get room_id error")
	}
	roomInfo := &RoomInfo{RoomId:roomId}
	err := roomInfo.GetRoomInfoByRoomId(roomId)
	if err != nil{
		return err
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &Client{
		hub: hub,
		conn: conn,
		send: make(chan []byte, 256),
		RoomInfo: roomInfo}
	roomInfoExpandFunc(roomInfo)
	clientExpandFunc(client)
	client.hub.register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
	return nil
}
