package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWTClaims struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var whitelist []string

func isWhitelisted(domain string, whitelist []string) bool {
	for _, item := range whitelist {
		if item == domain {
			return true
		}

	}
	return false
}

func init() {
	// 在 init 函数中初始化白名单列表
	whitelist = []string{
		"/auth/login",
		"/trading-day",
	}
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		if isWhitelisted(c.Request.URL.Path, whitelist) {
			return
		}

		// 从请求头中获取jwt-token
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		//解析JWT令牌,JWT 令牌字符串解析成一个结构体，以便访问其中的声明数据。在解析过程中，JWT 库会验证令牌的格式和签名是否正确，
		//但并不会验证令牌是否有效或令牌所包含的声明是否符合要求。解析 JWT 令牌的目的是将令牌中的声明数据提取出来，以便进行后续的处理和验证。
		token, err := jwt.ParseWithClaims(authorization, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 指定JWT签名秘钥
			return []byte(""), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			}
			c.Abort()
			return
		}

		// 验证JWT令牌
		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Set("userid", claims.UserId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
