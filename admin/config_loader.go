package admin

type ConfigLoaderName string

const (
	ConfigLoaderNameHTTP ConfigLoaderName = "http"
)

type ConfigLoader interface {
	Name() ConfigLoaderName
}

type BaseConfigLoader struct {
	LoaderName ConfigLoaderName `json:"module,omitempty"`
}

func (b *BaseConfigLoader) Name(name ConfigLoaderName) ConfigLoaderName {
	b.LoaderName = name
	return b.LoaderName
}

func NewConfigLoader[C ConfigLoader](configLoader C) C {
	configLoader.Name()
	return configLoader
}

type Header map[string][]string

type TLS struct {
	UseServerIdentity        bool     `json:"use_server_identity,omitempty"`
	ClientCertificateFile    string   `json:"client_certificate_file,omitempty"`
	ClientCertificateKeyFile string   `json:"client_certificate_key_file,omitempty"`
	RootCAPemFiles           []string `json:"root_ca_pem_files,omitempty"`
}

type ConfigLoaderHTTP struct {
	BaseConfigLoader
	Method  string `json:"method,omitempty"`
	URL     string `json:"url,omitempty"`
	Header  Header `json:"header,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
	TLS     *TLS   `json:"tls,omitempty"`
}

func (c *ConfigLoaderHTTP) Name() ConfigLoaderName {
	return c.BaseConfigLoader.Name(ConfigLoaderNameHTTP)
}

// interface guards
var (
	_ ConfigLoader = NewConfigLoader(&ConfigLoaderHTTP{})

	_ ConfigLoader = (*ConfigLoaderHTTP)(nil)
)
