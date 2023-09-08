package chttp

import (
	"github.com/aosasona/fate/pkg/candy/apps/chttp/providers"
)

// This package needs to be named like this to avoid a name collision with the net/http package

var (
	ListenerWrapperTLS      = tlsListenerWrapper{Wrapper: "tls"}
	ListenerWrapperRedirect = tlsListenerWrapper{Wrapper: "http_redirect"}
)

type ListenerWrapper interface {
	Listener() string
}

type Route struct {
	Group    string         `json:"group,omitempty"`
	Match    []RouteMatcher `json:"match,omitempty"`
	Handle   []RouteHandler `json:"handle,omitempty"`
	Terminal bool           `json:"terminal,omitempty"`
}

type Errors struct {
	Routes []Route `json:"routes"`
}

type tlsListenerWrapper struct {
	Wrapper string `json:"wrapper,omitempty"`
}

func (t tlsListenerWrapper) Listener() string { return t.Wrapper }

type AutomaticHTTPS struct {
	Disable                  bool     `json:"disable"`
	DisableRedirects         bool     `json:"disable_redirects"`
	DisableCertificates      bool     `json:"disable_certificates"`
	Skip                     []string `json:"skip"`
	SkipCertificates         []string `json:"skip_certificates"`
	IgnoreLoadedCertificates bool     `json:"ignore_loaded_certificates"`
}

type Logs struct {
	DefaultLoggerName    string            `json:"default_logger_name"`
	LoggerNames          map[string]string `json:"logger_names"`
	SkipHosts            []string          `json:"skip_hosts"`
	SkipUnmappedHosts    bool              `json:"skip_unmapped_hosts"`
	ShouldLogCredentials bool              `json:"should_log_credentials"`
}

type HTTP struct {
	HTTPPort      int               `json:"http_port,omitempty"`
	HTTPSPort     int               `json:"https_port,omitempty"`
	GracePeriod   int               `json:"grace_period,omitempty"`
	ShutDownDelay int               `json:"shutdown_delay,omitempty"`
	Servers       map[string]Server `json:"servers,omitempty"`
}

type Server struct {
	ID                    string                 `json:"@id,omitempty"`
	Listen                []string               `json:"listen,omitempty"`
	ListenerWrappers      []ListenerWrapper      `json:"listener_wrappers,omitempty"`
	ReadTimeout           int                    `json:"read_timeout,omitempty"`
	ReadHeaderTimeout     int                    `json:"read_header_timeout,omitempty"`
	WriteTimeout          int                    `json:"write_timeout,omitempty"`
	IdleTimeout           int                    `json:"idle_timeout,omitempty"`
	KeepAliveInterval     int                    `json:"keep_alive_interval,omitempty"`
	MaxHeaderBytes        int                    `json:"max_header_bytes,omitempty"`
	Routes                []Route                `json:"routes,omitempty"`
	Errors                *Errors                `json:"errors,omitempty"`
	NamedRoutes           map[string]Route       `json:"named_routes,omitempty"`
	TLSConnectionPolicies []TLSConnectionPolicy  `json:"tls_connection_policies,omitempty"`
	AutomaticHTTPS        *AutomaticHTTPS        `json:"automatic_https,omitempty"`
	StrictSNIHost         bool                   `json:"strict_sni_host,omitempty"`
	TrustedProxies        providers.TrustedProxy `json:"trusted_proxies,omitempty"`
	Logs                  *Logs                  `json:"logs,omitempty"`
	Protocols             []string               `json:"protocols,omitempty"`
	Metrics               any                    `json:"metrics,omitempty"` // experimental and subject to change, currently an unknown type
}
