package candy

import (
	"fmt"
	"net/http"

	"github.com/aosasona/fate/pkg/candy/apps/chttp"
)

type BaseMakerOpts struct {
	AppName         string
	Domain          string
	TargetPort      string
	HTTPListenPort  int
	HTTPSListenPort int
}

type RedirectRouteOpts struct {
	BaseMakerOpts
	TargetAddr string
}

func MakeRedirectRoute(opts RedirectRouteOpts) chttp.Server {
	listenPort := getPort(opts.HTTPListenPort, opts.HTTPSListenPort, false)

	return chttp.Server{
		Listen: []string{fmt.Sprintf(":%d", listenPort)},
		ID:     opts.AppName,
		Routes: []chttp.Route{
			{
				Match: []chttp.RouteMatcher{{chttp.MatchNameHost: chttp.MatchHost{opts.Domain}}},
				Handle: []chttp.RouteHandler{
					chttp.NewHandler(&chttp.RHStaticResponse{
						Headers:    chttp.StaticResponseHeaders{"Location": []string{"https://{http.request.host}{http.request.uri}"}},
						StatusCode: http.StatusPermanentRedirect,
					}),
				},
				Terminal: true,
			},
		},
	}
}

type ReverseProxyRouteOpts struct {
	BaseMakerOpts
	IsTLS      bool
	TargetAddr string
}

func MakeReverseProxyRoutes(opts ReverseProxyRouteOpts) chttp.Server {
	listenPort := getPort(opts.HTTPListenPort, opts.HTTPSListenPort, opts.IsTLS)

	return chttp.Server{
		Listen: []string{fmt.Sprintf(":%d", listenPort)},
		ID:     opts.AppName,
		Routes: []chttp.Route{
			{
				Match: []chttp.RouteMatcher{
					{
						chttp.MatchNameHost: chttp.MatchHost{opts.Domain},
					},
				},
				Handle: []chttp.RouteHandler{
					chttp.NewHandler(&chttp.RHSubRoute{
						Routes: []chttp.Route{
							{
								Handle: []chttp.RouteHandler{
									chttp.NewHandler(&chttp.RHReverseProxy{
										Upstreams: []chttp.Upstreams{
											{Dial: fmt.Sprintf("%s:%s", opts.TargetAddr, opts.TargetPort)},
										},
									}),
								},
							},
						},
					}),
				},
				Terminal: true,
			},
		},
	}
}

func getPort(httpPort, httpsPort int, isTLS bool) int {
	if isTLS {
		if httpsPort != 0 {
			return httpsPort
		}
		return 443
	}

	if httpPort != 0 {
		return httpPort
	}
	return 80
}
