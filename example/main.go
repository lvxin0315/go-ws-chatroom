package main

import (
	"github.com/lvxin0315/go-ws-chatroom"
	"net/http"
	"log"
)


func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func serveHome1(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home1.html")
}

func serveHome2(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home2.html")
}


func main()  {

	hub := GoWsChatroom.NewHub()
	go hub.Run()

	//TODO BEGIN 调试数据
	room1 := &GoWsChatroom.RoomInfo{RoomId: "123456-0"}
	room2 := &GoWsChatroom.RoomInfo{RoomId: "123456-1"}
	room3 := &GoWsChatroom.RoomInfo{RoomId: "123456-2"}
	room1.CreateRoom()
	room2.CreateRoom()
	room3.CreateRoom()
	//TODO err 应该显示房间已经存在
	err := room3.CreateRoom()
	log.Println(err)
	//TODO END


	http.HandleFunc("/h1", serveHome)
	http.HandleFunc("/h2", serveHome1)
	http.HandleFunc("/h3", serveHome2)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		GoWsChatroom.ServeWs(hub, w, r, func(roomInfo *GoWsChatroom.RoomInfo) {
			roomInfo.RoomInfoExpand = GoWsChatroom.RoomInfoExpand{}

		}, func(client *GoWsChatroom.Client) {
			client.ClientExpand = GoWsChatroom.ClientExpand{}
		})
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}
