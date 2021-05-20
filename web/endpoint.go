package web

import "golang.org/x/oauth2"

var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discordapp.com/api/oauth2/authorize",
	TokenURL:  "https://discordapp.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}
