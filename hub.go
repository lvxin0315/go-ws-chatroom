// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package GoWsChatroom

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *messageData

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	//房间客户端
	roomClients map[string]map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		broadcast:   make(chan *messageData),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		roomClients: make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.register:
			if h.roomClients[client.RoomInfo.RoomId] == nil {
				h.roomClients[client.RoomInfo.RoomId] = make(map[*Client]bool)
			}
			h.roomClients[client.RoomInfo.RoomId][client] = true

		case client := <-h.unregister:
			if _, ok := h.roomClients[client.RoomInfo.RoomId][client]; ok {
				delete(h.roomClients[client.RoomInfo.RoomId], client)
				close(client.send)
			}

		case messageData := <-h.broadcast:
			//房间信息解析
			for client := range h.roomClients[messageData.fromUser.RoomInfo.RoomId] {
				select {
				case client.send <- messageData.content:
				default:
					close(client.send)
					delete(h.roomClients[messageData.fromUser.RoomInfo.RoomId], client)
				}
			}
		}
	}
}
