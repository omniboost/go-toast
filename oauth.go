package asperion

import (
	"time"

	"golang.org/x/oauth2"
)

const (
	scope                = "*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.lsk.lightspeed.app/oauth/token",
			},
		},
	}

	return config
}
