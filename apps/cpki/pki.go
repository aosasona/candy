package cpki

import (
	"github.com/aosasona/fate/pkg/candy/storage"
)

type CARoot struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key,omitempty"`
	Format      string `json:"format,omitempty"`
}

// yes, they are the same thing but they could be different in the future and I don't want to change the name
type CAIntermediate struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key,omitempty"`
	Format      string `json:"format,omitempty"`
}

type CertificateAuthority struct {
	Name                   string          `json:"name"`
	RootCommonName         string          `json:"root_common_name,omitempty"`
	IntermediateCommonName string          `json:"intermediate_common_name,omitempty"`
	IntermidiateLifetime   string          `json:"intermediate_lifetime,omitempty"`
	InstallTrust           bool            `json:"install_trust,omitempty"`
	Root                   CARoot          `json:"root,omitempty"`
	Intermediate           CAIntermediate  `json:"intermediate,omitempty"`
	Storage                storage.Storage `json:"storage,omitempty"`
}

type CertificateAuthorities map[string]CertificateAuthority

type PKI struct {
	CertificateAuthorities CertificateAuthorities `json:"certificate_authorities"`
}
