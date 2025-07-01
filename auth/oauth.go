package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Client struct {
	config *oauth2.Config
}

var App Client

func Init_auth() {

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8081/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	App.config = config

}

func (a *Client) CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	token, err := a.config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(err)
	}
	client := a.config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	output, _ := io.ReadAll(response.Body)
	client_json := make(map[string]string)
	json.Unmarshal(output, &client_json)
	fmt.Println(client_json["email"])
	jwt_token, token_err := Create_JWT(client_json["email"])
	if token_err != nil {
		fmt.Println(token_err)
	}
	fmt.Println(jwt_token)
	fmt.Println(Verify_JWT(jwt_token))
	c.SetCookie("jwt", jwt_token, 3600, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/newApiKey")
}

func (a *Client) LoginHandler(c *gin.Context) {

	url := a.config.AuthCodeURL("random")
	c.Redirect(http.StatusTemporaryRedirect, url)
}
