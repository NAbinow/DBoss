package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAndVerifyCookies(c *gin.Context) (string, bool) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("No cookie found — redirecting to login.")
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return "", false
	}

	value, err := Verify_JWT(cookie)
	if err != nil {
		fmt.Println("JWT verification failed:", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return "", false
	}

	return value, true
}
