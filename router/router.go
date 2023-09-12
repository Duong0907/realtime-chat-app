package router

import (
	"go-chat/internal/middleware"
	"go-chat/internal/user"
	"go-chat/internal/ws"
	"log"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
func InitRouter(userPresentation *user.Presentation, wsPresentation *ws.Presentation) {
	r = gin.Default()
	r.Use(middleware.CorsMiddleware)

	r.POST("/api/user/signup", userPresentation.Handler.CreateUser)
	r.POST("/api/user/login", userPresentation.Handler.Login)
	r.GET("/api/user/logout", userPresentation.Handler.Logout)

	r.POST("/api/ws/createRoom", wsPresentation.Handler.CreateRoom)
	r.GET("/api/ws/joinRoom/:roomId", wsPresentation.Handler.JoinRoom)
	r.GET("/api/ws/getRooms", wsPresentation.Handler.GetRooms)
	r.GET("/api/ws/getClients/:roomId", wsPresentation.Handler.GetClients)

	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*.html")

	r.GET("/login", userPresentation.RenderLoginPage)
	r.GET("/signup", userPresentation.RenderSignupPage)
	r.GET("/", wsPresentation.RenderRoomPage)
	r.GET("/chat", wsPresentation.RenderChatPage)
}

func Start(addr string) error {
	log.Printf("Server is running at %s", addr)
	return r.Run(addr)
}
