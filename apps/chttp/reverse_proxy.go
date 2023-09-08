package chttp

type LoadBalancing struct {
	SelectionPolicy SelectionPolicy `json:"selection_policy,omitempty"`
	Retries         int             `json:"retries,omitempty"`
	TryDuration     int             `json:"try_duration,omitempty"`
	TryInterval     int             `json:"try_interval,omitempty"`
	RetryMatch      RouteMatcher    `json:"retry_match,omitempty"`
}

type (
	ActiveHealthCheck struct {
		Path         string              `json:"path,omitempty"`
		URI          string              `json:"uri,omitempty"`
		Port         int                 `json:"port,omitempty"`
		Headers      map[string][]string `json:"headers,omitempty"`
		Interval     int                 `json:"interval,omitempty"`
		Timeout      int                 `json:"timeout,omitempty"`
		MaxSize      int                 `json:"max_size,omitempty"`
		ExpectStatus int                 `json:"expect_status,omitempty"`
		ExpectBody   string              `json:"expect_body,omitempty"`
	}

	PassiveHealthCheck struct {
		FailDuration          int   `json:"fail_duration,omitempty"`
		MaxFails              int   `json:"max_fails,omitempty"`
		UnhealthyRequestCount int   `json:"unhealthy_request_count,omitempty"`
		UnhealthyStatus       []int `json:"unhealthy_status,omitempty"`
		UnhealthyLatency      int   `json:"unhealthy_latency,omitempty"`
	}

	HealthChecks struct {
		Active  *ActiveHealthCheck  `json:"active,omitempty"`
		Passive *PassiveHealthCheck `json:"passive,omitempty"`
	}
)

type Upstreams struct {
	Dial        string `json:"dial,omitempty"`
	LookupSrv   string `json:"lookup_srv,omitempty"`
	MaxRequests int    `json:"max_requests,omitempty"`
}

type HandlerHeader struct {
	HeaderRequest  `json:"header_request,omitempty"`
	HeaderResponse `json:"header_response,omitempty"`
}

type URISubstring struct {
	Find    string `json:"find,omitempty"`
	Replace string `json:"replace,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

type PathRegexp struct {
	Find    string `json:"find,omitempty"`
	Replace string `json:"replace,omitempty"`
}

type HandleResponseMatch struct {
	StatusCode []int               `json:"status_code,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
}

type HandleResponse struct {
	Match      *HandleResponseMatch `json:"match,omitempty"`
	StatusCode int                  `json:"status_code,omitempty"`
	Route      *Route               `json:"route,omitempty"`
}

type Rewrite struct {
	Method          string          `json:"method,omitempty"`
	URI             string          `json:"uri,omitempty"`
	StripPathPrefix string          `json:"strip_path_prefix,omitempty"`
	StripPathSuffix string          `json:"strip_path_suffix,omitempty"`
	URISubstring    []*URISubstring `json:"uri_substring,omitempty"`
	PathRegexp      []*PathRegexp   `json:"path_regexp,omitempty"`
}

type RHReverseProxy struct {
	BaseHandler
	Transport        Transport        `json:"transport,omitempty"`
	CircuitBeaker    any              `json:"circuit_beaker,omitempty"`
	LoadBalancing    *LoadBalancing   `json:"load_balancing,omitempty"`
	HealthChecks     *HealthChecks    `json:"health_checks,omitempty"`
	Upstreams        []Upstreams      `json:"upstreams,omitempty"`
	DynamicUpstreams DynamicUpstream  `json:"dynamic_upstreams,omitempty"`
	FlushInterval    int              `json:"flush_interval,omitempty"`
	TrustedProxies   []string         `json:"trusted_proxies,omitempty"`
	Headers          *HandlerHeader   `json:"headers,omitempty"`
	BufferRequests   bool             `json:"buffer_requests,omitempty"`
	BufferResponses  bool             `json:"buffer_responses,omitempty"`
	MaxBufferSize    int              `json:"max_buffer_size,omitempty"`
	RequestBuffers   int              `json:"request_buffers,omitempty"`
	ResponseBuffers  int              `json:"response_buffers,omitempty"`
	Rewrite          *Rewrite         `json:"rewrite,omitempty"`
	HandleResponse   []HandleResponse `json:"handle_response,omitempty"`
}
