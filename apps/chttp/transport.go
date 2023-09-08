package chttp

type ProtocolName string

const (
	ProtocolHTTP    ProtocolName = "http"
	ProtocolFastCGI ProtocolName = "fastcgi"
)

type Transport interface {
	Transport() ProtocolName
}

type BaseTransport struct {
	Protocol ProtocolName `json:"protocol"`
}

type (
	Resolver struct {
		Addresses []string `json:"addresses"`
	}

	TLS struct {
		RootCAPool                []string `json:"root_ca_pool,omitempty"`
		RootCAPemFiles            []string `json:"root_ca_pem_files,omitempty"`
		ClientCertificateFile     string   `json:"client_certificate_file,omitempty"`
		ClientCertificateKeyFile  string   `json:"client_certificate_key_file,omitempty"`
		ClientCertificateAutomate string   `json:"client_certificate_automate,omitempty"`
		InsecureSkipVerify        bool     `json:"insecure_skip_verify,omitempty"`
		HandshakeTimeout          int      `json:"handshake_timeout,omitempty"`
		ServerName                string   `json:"server_name,omitempty"`
		Renegotiation             string   `json:"renegotiation"`
		ExceptPorts               []int    `json:"except_ports"`
	}

	KeepAlive struct {
		Enabled             bool `json:"enabled"`
		ProbeInterval       int  `json:"probe_interval,omitempty"`
		MaxIdleConns        int  `json:"max_idle_conns,omitempty"`
		MaxIdleConnsPerHost int  `json:"max_idle_conns_per_host,omitempty"`
		IdleTimeout         int  `json:"idle_timeout,omitempty"`
	}

	TransportHTTP struct {
		BaseTransport
		Resolver              Resolver  `json:"resolver,omitempty"`
		TLS                   TLS       `json:"tls,omitempty"`
		KeepAlive             KeepAlive `json:"keep_alive,omitempty"`
		Compression           bool      `json:"compression,omitempty"`
		MaxConnsPerHost       int       `json:"max_conns_per_host,omitempty"`
		DialTimeout           int       `json:"dial_timeout,omitempty"`
		DialFallbackDelay     int       `json:"dial_fallback_delay,omitempty"`
		ResponseHeaderTimeout int       `json:"response_header_timeout,omitempty"`
		ExpectContinueTimeout int       `json:"expect_continue_timeout,omitempty"`
		MaxResponseHeaderSize int       `json:"max_response_header_size,omitempty"`
		WriteBufferSize       int       `json:"write_buffer_size,omitempty"`
		ReadBufferSize        int       `json:"read_buffer_size,omitempty"`
		ReadTimeout           int       `json:"read_timeout,omitempty"`
		WriteTimeout          int       `json:"write_timeout,omitempty"`
		Versions              []string  `json:"versions,omitempty"`
	}
)

type TransportFastCGI struct {
	BaseTransport
	Root               string            `json:"root,omitempty"`
	SplitPath          []string          `json:"split_path,omitempty"`
	ResolveRootSymlink bool              `json:"resolve_root_symlink,omitempty"`
	Env                map[string]string `json:"env,omitempty"`
	DialTimeout        int               `json:"dial_timeout,omitempty"`
	ReadTimeout        int               `json:"read_timeout,omitempty"`
	WriteTimeout       int               `json:"write_timeout,omitempty"`
	CaptureStderr      bool              `json:"capture_stderr"`
}

func (bt *BaseTransport) Transport(name ProtocolName) ProtocolName {
	bt.Protocol = name
	return bt.Protocol
}

func NewTransport[T Transport](t T) T {
	t.Transport()
	return t
}

func (t *TransportHTTP) Transport() ProtocolName { return t.BaseTransport.Transport(ProtocolHTTP) }

func (t *TransportFastCGI) Transport() ProtocolName {
	return t.BaseTransport.Transport(ProtocolFastCGI)
}

// Interface guards
var (
	_ Transport = (*TransportHTTP)(nil)
	_ Transport = (*TransportFastCGI)(nil)
)
