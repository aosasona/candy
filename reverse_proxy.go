package candy

import (
	"fmt"

	"github.com/aosasona/fate/internal/config"
	"github.com/aosasona/fate/pkg/candy/apps"
	"github.com/aosasona/fate/pkg/candy/apps/chttp"
	"github.com/aosasona/fate/pkg/candy/apps/ctls"
	"github.com/aosasona/fate/pkg/candy/apps/ctls/issuers"
	"github.com/aosasona/fate/pkg/candy/internal/request"
)

type ReverseProxyOpts struct {
	AppName    string
	Domain     string
	Port       string
	TargetAddr string
}

func CreateReverseProxy(opts ReverseProxyOpts) error {
	tlsAppName := fmt.Sprintf("%s_tls", opts.AppName)

	var (
		certIssuers ctls.Issuers
		onDemand    bool
	)

	if config.IsDev {
		certIssuers = ctls.Issuers{issuers.NewIssuer(&issuers.IssuerInternal{})}
	} else {
		certIssuers = ctls.Issuers{
			issuers.NewIssuer(&issuers.IssuerAcme{Email: config.AcmeEmail}),
			issuers.NewIssuer(&issuers.IssuerZeroSSL{Email: config.AcmeEmail}),
		}
		onDemand = true
	}

	req := apps.App{
		HTTP: &chttp.HTTP{
			Servers: map[string]chttp.Server{
				opts.AppName: MakeReverseProxyRoutes(
					ReverseProxyRouteOpts{
						BaseMakerOpts: BaseMakerOpts{AppName: opts.AppName, Domain: opts.Domain, TargetPort: opts.Port},
						TargetAddr:    opts.TargetAddr,
					},
				),
				tlsAppName: MakeReverseProxyRoutes(
					ReverseProxyRouteOpts{
						BaseMakerOpts: BaseMakerOpts{AppName: tlsAppName, Domain: opts.Domain, TargetPort: opts.Port},
						TargetAddr:    opts.TargetAddr,
						IsTLS:         true,
					},
				),
			},
		},

		TLS: &ctls.TLS{
			Automation: &ctls.Automation{
				Policies: []ctls.Policy{
					{Subjects: []string{opts.Domain}, Issuers: certIssuers, OnDemand: onDemand},
				},
			},
		},
	}

	var response map[string]interface{}
	err := request.Post(request.Request{Path: "/config/apps", Data: req}, &response)
	if err != nil {
		return err
	}

	return nil
}
