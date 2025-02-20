package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *CartController) CheckSessionID(ctx *gin.Context) {
	var sessionID string
	sessionIDCookie, err := ctx.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		sessionID = time.Now().String()
		ctx.SetCookie("ice_session_id", sessionID, 3600, "/", "localhost", false, true)
	} else {
		sessionID = sessionIDCookie.Value
	}

	ctx.Set(sessionIDKey, sessionID)
}
