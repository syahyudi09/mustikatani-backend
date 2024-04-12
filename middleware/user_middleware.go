package middleware

import (
	"net/http"
	"pinjam-modal-app/utils/authutil"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check exist token
		h := &authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		// check token kosong
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		// check verify token
		token, err := authutil.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		username := session.Get("Username")
		userRole := session.Get("UserRole")
		if userRole == nil && username == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}
		if userRole == "Customer" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}