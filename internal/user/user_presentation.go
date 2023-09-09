package user

import (
	"net/http"
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

func (p *Presentation) RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (p *Presentation) RenderSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}