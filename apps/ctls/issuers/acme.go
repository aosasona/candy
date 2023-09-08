package issuers

type ExternalAccount struct {
	KeyID  string `json:"key_id"`
	MacKey string `json:"mac_key"`
}

type ChallengeHTTP struct {
	Disabled      bool `json:"disabled,omitempty"`
	AlternatePort int  `json:"alternate_port,omitempty"`
}

type ChallengeTLSALPN struct {
	Disabled      bool `json:"disabled,omitempty"`
	AlternatePort int  `json:"alternate_port,omitempty"`
}

type (
	DNSProviderName string

	DNSProvider interface {
		Name() DNSProviderName
	}
)

type ChallengeDNS struct {
	Provider           DNSProvider `json:"provider,omitempty"`
	TTL                int         `json:"ttl,omitempty"`
	PropagationDelay   int         `json:"propagation_delay,omitempty"`
	PropagationTimeout int         `json:"propagation_timeout,omitempty"`
	Resolvers          []string    `json:"resolvers,omitempty"`
	OverrideDomain     string      `json:"override_domain,omitempty"`
}

type Challenges struct {
	HTTP     *ChallengeHTTP    `json:"http,omitempty"`
	TLSALPN  *ChallengeTLSALPN `json:"tls-alpn,omitempty"`
	DNS      *ChallengeDNS     `json:"dns,omitempty"`
	BindHost string            `json:"bind_host,omitempty"`
}

type PreferredChains struct {
	Smallest       bool   `json:"smallest,omitempty"`
	RootCommonName string `json:"root_common_name,omitempty"`
	AnyCommonName  string `json:"any_common_name,omitempty"`
}

type IssuerAcme struct {
	BaseIssuer
	CA                   string           `json:"ca,omitempty"`
	TestCA               string           `json:"test_ca,omitempty"`
	Email                string           `json:"email,omitempty"`
	AccountKey           string           `json:"account_key,omitempty"`
	ExternalAccount      *ExternalAccount `json:"external_account,omitempty"`
	AcmeTimeout          int              `json:"acme_timeout,omitempty"`
	Challenges           *Challenges      `json:"challenges,omitempty"`
	TrustedRootsPemFiles []string         `json:"trusted_roots_pem_files,omitempty"`
	PreferredChains      *PreferredChains `json:"preferred_chains,omitempty"`
}

func (i *IssuerAcme) Name() IssuerName { return i.BaseIssuer.Name(IssuerNameACME) }
