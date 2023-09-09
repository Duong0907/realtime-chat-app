package ws

import (
	"log"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			log.Println("A new user registers")
			// If there is a room of that client
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				// If the client is not in the room, add it
				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
					log.Println("A new user register")
				}
			}
		case cl := <-h.Unregister:
			// If there is a room of that client
			if _, ok := h.Rooms[cl.RoomID]; ok {
				// If that room have that client in it
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						log.Println("Unregister")
						h.Broadcast <- &Message{
							Content:  "A user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
