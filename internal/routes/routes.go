package routes

import (
	"fmt"
	"net/http"

	"github.com/evanjo03/oauth/internal/auth0"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func OauthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handleError(w, fmt.Errorf("invalid request method"))
		return
	}

	if err := r.ParseForm(); err != nil {
		handleError(w, err)
		return
	}

	domain := r.FormValue("domain")
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	username := r.FormValue("username")
	password := r.FormValue("password")
	audience := r.FormValue("audience")
	scope := r.FormValue("scope")
	grantType := r.FormValue("grant_type")

	request := auth0.Request{
		GrantType:    grantType,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		Audience:     audience,
		Scope:        scope,
	}

	accessToken, err := auth0.GetToken(request, domain)
	if err != nil {
		handleError(w, err)
	}

	if accessToken == "" {
		handleError(w, fmt.Errorf("access token is empty"))
		return
	}

	fmt.Fprint(w, accessToken)
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "Error: %v", err)
}
