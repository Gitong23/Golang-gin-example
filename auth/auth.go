package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// func AccessToken(c *gin.Context) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
// 		Issuer:    "skooldio",
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
// 	})
// 	ss, err := token.SignedString([]byte("==signature=="))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": ss,
// 	})

// }

//
// func AccessToken(c *gin.Context) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
// 		Issuer:    "skooldio",
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
// 	})
// 	ss, err := token.SignedString([]byte("==signature=="))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": ss,
// 	})

// }

//refactor higher order func
func AccessToken(signature string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			Issuer:    "skooldio",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		})
		ss, err := token.SignedString([]byte(signature))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": ss,
		})
	}
}
