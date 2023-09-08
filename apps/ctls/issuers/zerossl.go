package issuers

type IssuerZeroSSL struct {
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
	APIKey               string           `json:"api_key,omitempty"`
}

func (i *IssuerZeroSSL) Name() IssuerName { return i.BaseIssuer.Name(IssuerNameZeroSSL) }
