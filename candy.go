package candy

import (
	"net/http"

	"github.com/aosasona/fate/pkg/candy/admin"
	"github.com/aosasona/fate/pkg/candy/apps"
	"github.com/aosasona/fate/pkg/candy/internal/client"
	"github.com/aosasona/fate/pkg/candy/internal/request"
	"github.com/aosasona/fate/pkg/candy/storage"
)

/*
`candy` is a package to interact with the caddy REST API

This is a wrapper around the Caddy API, it may not contain all the features of the Caddy API, or the full JSON structure, I will add them as I need them.

The Caddy API is documented here: https://caddyserver.com/docs/api-tutorial

* Why did I do this?

- Could I have just used maps? Yes

- Could I have just used Caddyfiles as templates? also yes

I mostly did this because I wanted to, it will eventually live it it's own repo, but for now it's here. I created the repo months ago, never got around to it and this was the perfect chance to work on it.
*/

type CaddyRequest struct {
	Apps    apps.App        `json:"apps,omitempty"`
	Storage storage.Storage `json:"storage,omitempty"`
	Admin   admin.Admin     `json:"admin,omitempty"`
}

func init() {
	client.Client = &client.Candy{
		HttpClient: &http.Client{},
	}
}

var EnableDebug = request.EnableDebugMode

func SetCaddyAddr(url string) {
	request.CaddyAPIAddr = url
}
