package chttp

type DynamicUpstreamName string

const (
	DynamicUpstreamNameA     DynamicUpstreamName = "a"
	DynamicUpstreamNameMulti DynamicUpstreamName = "multi"
	DynamicUpstreamNameSRV   DynamicUpstreamName = "srv"
)

type DynamicUpstream interface {
	Upstream() DynamicUpstreamName
}

type BaseDynamicUpstream struct {
	Source string `json:"source"`
}

func (b *BaseDynamicUpstream) Upstream(name DynamicUpstreamName) DynamicUpstreamName {
	b.Source = string(name)
	return name
}

func NewDynamicUpstream[U DynamicUpstream](upstream U) U {
	upstream.Upstream()
	return upstream
}

// DynamicUpstream types definitions
type DynamicUpstreamA struct {
	BaseDynamicUpstream
	Name              string   `json:"name"`
	Port              string   `json:"port"`
	Refresh           int      `json:"refresh"`
	Resolver          Resolver `json:"resolver"`
	DialTimeout       int      `json:"dial_timeout"`
	DialFallbackDelay int      `json:"dial_fallback_delay"`
}

type DynamicUpstreamMulti struct {
	BaseDynamicUpstream
	Sources []DynamicUpstream `json:"sources"`
}

type DynamicUpstreamSRV struct {
	BaseDynamicUpstream
	Service           string   `json:"service"`
	Proto             string   `json:"proto"`
	Name              string   `json:"name"`
	Refresh           int      `json:"refresh"`
	Resolver          Resolver `json:"resolver"`
	DialTimeout       int      `json:"dial_timeout"`
	DialFallbackDelay int      `json:"dial_fallback_delay"`
}

// Methods
func (d *DynamicUpstreamA) Upstream() DynamicUpstreamName {
	return d.BaseDynamicUpstream.Upstream(DynamicUpstreamNameA)
}

func (d *DynamicUpstreamMulti) Upstream() DynamicUpstreamName {
	return d.BaseDynamicUpstream.Upstream(DynamicUpstreamNameMulti)
}

func (d *DynamicUpstreamSRV) Upstream() DynamicUpstreamName {
	return d.BaseDynamicUpstream.Upstream(DynamicUpstreamNameSRV)
}

// Interface guards

var (
	_ DynamicUpstream = (*DynamicUpstreamA)(nil)
	_ DynamicUpstream = (*DynamicUpstreamMulti)(nil)
	_ DynamicUpstream = (*DynamicUpstreamSRV)(nil)
)
