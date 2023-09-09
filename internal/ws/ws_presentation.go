package ws

import (
	"log"
	"net/http"

	"go-chat/internal/authorization"

	"github.com/gin-gonic/gin"
)

type Presentation struct {
	Handler
}

func NewPresentation(h Handler) *Presentation {
	return &Presentation{
		Handler: h,
	}
}

func (h *Presentation) RenderRoomPage(c *gin.Context) {
	tokenSTring, err := c.Cookie("jwt")
	log.Println(tokenSTring)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		return
	}

	_, err = auth.ParseTokenString(tokenSTring)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		return
	}

	// Get rooms
	rooms := make([]RoomRes, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.HTML(http.StatusOK, "room.html", RoomData{Rooms: rooms})
}

type ChatInfo struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	RoomName string `json:"room_name"`
	RoomID   string `json:"room_id"`
}

func (h *Presentation) RenderChatPage(c *gin.Context) {
	tokenSTring, err := c.Cookie("jwt")
	log.Println(tokenSTring)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		return
	}

	claims, err := auth.ParseTokenString(tokenSTring)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		return
	}

	var info ChatInfo
	info.Username = claims.Username
	info.UserID = claims.ID
	info.RoomName = c.Query("room_name")
	info.RoomID = c.Query("room_id")

	log.Println(claims)

	c.HTML(http.StatusOK, "chat.html", info)
}
