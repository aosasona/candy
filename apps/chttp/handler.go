package chttp

import (
	"github.com/aosasona/fate/pkg/candy/apps/chttp/providers"
)

type HandlerName string

const (
	HandlerInvoke              HandlerName = "invoke"
	HandlerError               HandlerName = "error"
	HandlerStaticResponse      HandlerName = "static_response"
	HandlerSubRoute            HandlerName = "subroute"
	HandlerVars                HandlerName = "vars"
	HandlerAuthentication      HandlerName = "authentication"
	HandlerEncode              HandlerName = "encode"
	HandlerFileServer          HandlerName = "file_server"
	HandlerHeaders             HandlerName = "headers"
	HandlerMap                 HandlerName = "map"
	HandlerPush                HandlerName = "push"
	HandlerRequestBody         HandlerName = "request_body"
	HandlerCopyResponse        HandlerName = "copy_response"
	HandlerCopyResponseHeaders HandlerName = "copy_response_headers"
	HandlerReverseProxy        HandlerName = "reverse_proxy"
	HandlerRewrite             HandlerName = "rewrite"
	HandlerTemplates           HandlerName = "templates"
	HandlerTracing             HandlerName = "tracing"
	HandlerAcmeServer          HandlerName = "acme_server"
	HandlerMetrics             HandlerName = "metrics"
)

// NewHandler creates a new handler based on the provided handler with the correct handler name included
//
// Usage: candy.NewHandler(&candy.RHStaticResponse{}) You can also set fields on the handler, for example: candy.NewHandler(&candy.RHStaticResponse{StatusCode: 200})
func NewHandler[H RouteHandler](handler H) H {
	handler.Handler()
	return handler
}

type RouteHandler interface {
	Handler() HandlerName
}

type StaticResponseHeaders map[string][]string

type BaseHandler struct {
	HandlerName HandlerName `json:"handler"`
}

func (h *BaseHandler) Handler(name HandlerName) HandlerName {
	h.HandlerName = name
	return h.HandlerName
}

// RouteHandler types definitions
type RHInvoke struct {
	BaseHandler
	Name string `json:"name,omitempty"`
}

type RHError struct {
	BaseHandler
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
}

type RHStaticResponse struct {
	BaseHandler
	StatusCode int                   `json:"status_code,omitempty"`
	Headers    StaticResponseHeaders `json:"headers,omitempty"`
	Body       string                `json:"body,omitempty"`
	Close      bool                  `json:"close,omitempty"`
	Abort      bool                  `json:"abort,omitempty"`
}

type RHSubRoute struct {
	BaseHandler
	Routes []Route `json:"routes,omitempty"`
	Errors *Errors `json:"errors,omitempty"`
}

type RHVars struct {
	BaseHandler
}

type (
	Providers        map[providers.AuthProviderName]providers.AuthProvider
	RHAuthentication struct {
		BaseHandler
		Providers Providers `json:"providers,omitempty"`
	}
)

type (
	Encoders map[providers.EncodingName]providers.Encoding

	EncodingMatch struct {
		StatusCode int                 `json:"status_code,omitempty"`
		Headers    map[string][]string `json:"headers,omitempty"`
	}

	RHEncode struct {
		BaseHandler
		Encodings     Encoders      `json:"encodings,omitempty"`
		Prefer        []string      `json:"prefer,omitempty"`
		MinimumLength int           `json:"minimum_length,omitempty"`
		Match         EncodingMatch `json:"match,omitempty"`
	}
)

type RHFileServer struct {
	BaseHandler
	FileSystem any      `json:"file_system,omitempty"`
	Root       string   `json:"root,omitempty"`
	Hide       []string `json:"hide,omitempty"`
	IndexNames []string `json:"index_names,omitempty"`
	Browse     struct {
		TemplateFile string `json:"template_file,omitempty"`
	} `json:"browse,omitempty"`
	CanonicalURIs      bool     `json:"canonical_uris,omitempty"`
	StatusCode         int      `json:"status_code,omitempty"`
	PassThru           bool     `json:"pass_thru,omitempty"`
	Precompressed      Encoders `json:"precompressed,omitempty"`
	PrecompressedOrder []string `json:"precompressed_order,omitempty"`
}

type (
	HeaderRequestReplace struct {
		Search       string `json:"search"`
		SearchRegexp string `json:"search_regexp"`
		Replace      string `json:"replace"`
	}

	HeaderRequest struct {
		Add     map[string][]string               `json:"add,omitempty"`
		Set     map[string][]string               `json:"set,omitempty"`
		Delete  []string                          `json:"delete,omitempty"`
		Replace map[string][]HeaderRequestReplace `json:"replace,omitempty"`
	}

	HeaderResponseRequire struct {
		StatusCode []int               `json:"status_code,omitempty"`
		Headers    map[string][]string `json:"headers,omitempty"`
	}

	HeaderResponse struct {
		HeaderRequest
		Require  HeaderResponseRequire `json:"require,omitempty"`
		Deferred bool                  `json:"deferred,omitempty"`
	}

	RHHeaders struct {
		BaseHandler
		Request HeaderRequest `json:"request,omitempty"`
	}
)

type (
	Mapping struct {
		Input       string   `json:"input,omitempty"`
		InputRegexp string   `json:"input_regexp,omitempty"`
		Outputs     []string `json:"outputs,omitempty"`
	}

	RHMap struct {
		BaseHandler
		Source       string    `json:"source,omitempty"`
		Destinations []string  `json:"destinations,omitempty"`
		Mappings     []Mapping `json:"mappings,omitempty"`
		Defaults     []string  `json:"defaults,omitempty"`
	}
)

type (
	Resource struct {
		Method string `json:"method"`
		Target string `json:"target"`
	}

	RHPush struct {
		BaseHandler
		Headers   HeaderRequest `json:"headers"`
		Resources []Resource    `json:"resources"`
	}
)

type RHRequestBody struct {
	BaseHandler
	MaxSize int `json:"max_size"`
}

type RHCopyResponse struct {
	BaseHandler
	StatusCode int `json:"status_code"`
}

type RHCopyResponseHeaders struct {
	BaseHandler
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type RHRewrite struct {
	BaseHandler
	Rewrite
}

type RHTemplate struct {
	BaseHandler
	FileRoot   string   `json:"file_root"`
	MimeTypes  []string `json:"mime_types"`
	Delimiters []string `json:"delimiters"`
}

type RHTracing struct {
	BaseHandler
	Span string `json:"span"`
}

type RHAcmeServer struct {
	BaseHandler
	CA           string `json:"ca"`
	Lifetime     int    `json:"lifetime"`
	Host         string `json:"host"`
	PathPrefix   string `json:"path_prefix"`
	SignWithRoot bool   `json:"sign_with_root"`
}

type RHMetrics struct {
	BaseHandler
	DisableOpenmetrics bool `json:"disable_openmetrics"`
}

/** RouteHandler method definitions */
func (r *RHInvoke) Handler() HandlerName { return r.BaseHandler.Handler(HandlerInvoke) }

func (r *RHError) Handler() HandlerName { return r.BaseHandler.Handler(HandlerError) }

func (r *RHStaticResponse) Handler() HandlerName { return r.BaseHandler.Handler(HandlerStaticResponse) }

func (r *RHSubRoute) Handler() HandlerName { return r.BaseHandler.Handler(HandlerSubRoute) }

func (r *RHVars) Handler() HandlerName { return r.BaseHandler.Handler(HandlerVars) }

func (r *RHAuthentication) Handler() HandlerName { return r.BaseHandler.Handler(HandlerAuthentication) }

func (r *RHEncode) Handler() HandlerName { return r.BaseHandler.Handler(HandlerEncode) }

func (r *RHFileServer) Handler() HandlerName { return r.BaseHandler.Handler(HandlerFileServer) }

func (r *RHHeaders) Handler() HandlerName { return r.BaseHandler.Handler(HandlerHeaders) }

func (r *RHMap) Handler() HandlerName { return r.BaseHandler.Handler(HandlerMap) }

func (r *RHPush) Handler() HandlerName { return r.BaseHandler.Handler(HandlerPush) }

func (r *RHRequestBody) Handler() HandlerName { return r.BaseHandler.Handler(HandlerRequestBody) }

func (r *RHCopyResponse) Handler() HandlerName { return r.BaseHandler.Handler(HandlerCopyResponse) }

func (r *RHCopyResponseHeaders) Handler() HandlerName {
	return r.BaseHandler.Handler(HandlerCopyResponseHeaders)
}

func (r *RHReverseProxy) Handler() HandlerName {
	return r.BaseHandler.Handler(HandlerReverseProxy)
}
func (r *RHRewrite) Handler() HandlerName { return r.BaseHandler.Handler(HandlerRewrite) }

func (r *RHTemplate) Handler() HandlerName { return r.BaseHandler.Handler(HandlerTemplates) }

func (r *RHTracing) Handler() HandlerName { return r.BaseHandler.Handler(HandlerTracing) }

func (r *RHAcmeServer) Handler() HandlerName { return r.BaseHandler.Handler(HandlerAcmeServer) }

func (r *RHMetrics) Handler() HandlerName { return r.BaseHandler.Handler(HandlerMetrics) }

// Interface guards
var (
	_ RouteHandler = (*RHInvoke)(nil)
	_ RouteHandler = (*RHError)(nil)
	_ RouteHandler = (*RHStaticResponse)(nil)
	_ RouteHandler = (*RHSubRoute)(nil)
	_ RouteHandler = (*RHVars)(nil)
	_ RouteHandler = (*RHAuthentication)(nil)
	_ RouteHandler = (*RHEncode)(nil)
	_ RouteHandler = (*RHFileServer)(nil)
	_ RouteHandler = (*RHHeaders)(nil)
	_ RouteHandler = (*RHMap)(nil)
	_ RouteHandler = (*RHPush)(nil)
	_ RouteHandler = (*RHRequestBody)(nil)
	_ RouteHandler = (*RHCopyResponse)(nil)
	_ RouteHandler = (*RHCopyResponseHeaders)(nil)
	_ RouteHandler = (*RHReverseProxy)(nil)
	_ RouteHandler = (*RHRewrite)(nil)
	_ RouteHandler = (*RHTemplate)(nil)
	_ RouteHandler = (*RHTracing)(nil)
	_ RouteHandler = (*RHAcmeServer)(nil)
	_ RouteHandler = (*RHMetrics)(nil)
)
