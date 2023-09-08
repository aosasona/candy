package apps

import (
	"github.com/aosasona/fate/pkg/candy/apps/chttp"
	"github.com/aosasona/fate/pkg/candy/apps/cpki"
	"github.com/aosasona/fate/pkg/candy/apps/ctls"
)

type AppName string

const (
	HTTP AppName = "http"
	TLS  AppName = "tls"
	PKI  AppName = "pki"
)

type App struct {
	HTTP *chttp.HTTP `json:"http,omitempty"`
	TLS  *ctls.TLS   `json:"tls,omitempty"`
	PKI  *cpki.PKI   `json:"pki,omitempty"`
}
