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
		Broadcast:  make(chan *Message, 10),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// If there is a room of that client
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				// If the client is not in the room, add it
				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
					log.Printf("User %v joined the chat\n", cl.Username)
				}
			}

		case cl := <-h.Unregister:
			// If there is a room of that client
			if _, ok := h.Rooms[cl.RoomID]; ok {
				// If that room have that client in it
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)

					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "A user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					
					log.Printf("User %v left the chat\n", cl.Username)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					log.Println(cl.Username)
					cl.Message <- m
				}
			}
		}
	}
}
