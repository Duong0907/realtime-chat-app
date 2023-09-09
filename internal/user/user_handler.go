// INTERFACE ADAPTERS LAYER
// This layer holds:
// - Presenters (UI Logic, States)
// - Controllers (Interface that holds methods needed by the application
// which is implemented by Web, Devices or External Interfaces)
// - Gateways (Interface that holds every CRUD operation
// performed by the application, implemented by DB)

package user

import (
	// "log"
	"log"
	"net/http"
	"strings"

	
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service 
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq 
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(u.AccessToken)
	domain := c.Request.Host
	domain = strings.Split(domain, ":")[0]
	log.Println(domain)
	c.SetCookie("jwt", u.AccessToken, 3600, "/", domain, false, true)
	c.JSON(http.StatusOK, &LoginRes{
		ID: u.ID,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
}