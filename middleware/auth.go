package middleware

// import (
// 	"errors"
// 	"net/http"
// 	"strings"

// 	"pinjam-modal-app/model"
// 	"pinjam-modal-app/repository"
// 	"pinjam-modal-app/utils/authutil"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func AuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		func AuthenticationMiddleware() gin.HandlerFunc {
// 			return func(c *gin.Context) {
// 				// Periksa token dari header Authorization
// 				authorizationHeader := c.GetHeader("Authorization")
// 				if authorizationHeader == "" {
// 					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 					c.Abort()
// 					return
// 				}

// 				// Misalnya, format header Authorization: Bearer <token>
// 				splitToken := strings.Split(authorizationHeader, "Bearer ")
// 				if len(splitToken) != 2 {
// 					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 					c.Abort()
// 					return
// 				}

// 				tokenString := splitToken[1]

// Validasi token
// token, err := validateToken(tokenString)
// if err != nil {
// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 	c.Abort()
// 	return
// }

// // Periksa apakah pengguna telah logout
// if isTokenLoggedOut(tokenString) {
// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 	c.Abort()
// 	return
// }

// 				// Setel data pengguna pada konteks Gin
// 				user, err := getUserFromToken(token)
// 				if err != nil {
// 					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 					c.Abort()
// 					return
// 				}

// 				c.Set("user", user)

// 				// Periksa izin akses berdasarkan peran pengguna
// 				if !checkAccess(user.Role, c.Request.URL.Path) {
// 					c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
// 					c.Abort()
// 					return
// 				}

// 				c.Next()
// 			}
// 		}

// 		// Periksa izin akses berdasarkan peran pengguna
// 		user := c.MustGet("user").(*model.UserModel)
// 		if !checkAccess(user.Role, c.Request.URL.Path) {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func checkAccess(role string, path string) bool {
// 	// Implementasi logika untuk memeriksa izin akses berdasarkan peran pengguna
// 	// Misalnya, Anda dapat menggunakan peta (map) yang memetakan peran pengguna ke rute yang diizinkan

// 	return token, nil
// }

// func isTokenLoggedOut(tokenString string) bool {
// 	// Periksa apakah token telah logout menggunakan repositori LogoutRepo
// 	return userRepo.IsTokenLoggedOut(tokenString)
// }

// 	for _, route := range allowedRoutes {
// 		if route == path {
// 			return true
// 		}
// 	}

// 	return false
// }
