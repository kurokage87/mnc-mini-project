package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	controllers "mnc/controller"
	"net/http"
	"strings"
)

func Route() *gin.Engine {
	router := gin.Default()

	//router.Use(AllowedMethodsHandler())

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	authorized := router.Group("/", AuthMiddleware())
	{
		authorized.POST("/profile", controllers.UpdateProfile)
		authorized.POST("/topup", controllers.Topup)
		authorized.POST("/payment", controllers.Payment)
		authorized.POST("/transfer", controllers.Transfer)
		authorized.GET("/transactions", controllers.Report)
	}

	return router
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte("your_secret_key"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}
	}
}
