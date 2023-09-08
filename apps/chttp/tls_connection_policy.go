package chttp

import (
	"github.com/aosasona/fate/pkg/candy/apps/chttp/providers"
)

// HandshakeMatch
type HandshakeMatchName string

const (
	HandshakeMatchNameRemoteIP HandshakeMatchName = "remote_ip"
	HandshakeMatchNameSNI      HandshakeMatchName = "sni"
)

type HandshakeMatch interface {
	Name() HandshakeMatchName
}

// Built-in handshake matches structs
type HandshakeMatchRemoteIP struct {
	Ranges    []string `json:"ranges"`
	NotRanges []string `json:"not_ranges"`
}

type HandshakeMatchSNI []string

// Built-in handshake matches methods to implement HandshakeMatch interface
func (h HandshakeMatchRemoteIP) Name() HandshakeMatchName { return HandshakeMatchNameRemoteIP }

func (h HandshakeMatchSNI) Name() HandshakeMatchName { return HandshakeMatchNameSNI }

// Other deps
type CertificateSelection struct {
	SerialNumber         []string `json:"serial_number"`
	SubjectOrganiszation []string `json:"subject_organization"`
	PublicKeyAlgorithm   int      `json:"public_key_algorithm"`
	AnyTag               []string `json:"any_tag"`
	AllTags              []string `json:"all_tags"`
}

type ClientAuthentication struct {
	TrustedCACerts         []string                 `json:"trusted_ca_certs"`
	TrustedCACertsPemFiles []string                 `json:"trusted_ca_certs_pem_files"`
	TrustedLeafCerts       []string                 `json:"trusted_leaf_certs"`
	Verifiers              []providers.AuthVerifier `json:"verifiers"`
	Mode                   string                   `json:"mode"`
}

// TLSConnectionPolicy
type TLSConnectionPolicy struct {
	Match                map[HandshakeMatchName]HandshakeMatch `json:"match,omitempty"`
	CertificateSelection CertificateSelection                  `json:"certificate_selection,omitempty"`
	CipherSuites         []string                              `json:"cipher_suites,omitempty"`
	Curves               []string                              `json:"curves,omitempty"`
	Alpn                 []string                              `json:"alpn,omitempty"`
	ProtocolMin          string                                `json:"protocol_min,omitempty"`
	ProtocolMax          string                                `json:"protocol_max,omitempty"`
	ClientAuthentication ClientAuthentication                  `json:"client_authentication,omitempty"`
	DefaultSNI           string                                `json:"default_sni,omitempty"`
	FallbackSNI          string                                `json:"fallback_sni,omitempty"`
	InsecureSecretsLog   string                                `json:"insecure_secrets_log,omitempty"`
}

// Interface guards
var (
	_ HandshakeMatch = (*HandshakeMatchRemoteIP)(nil)
	_ HandshakeMatch = (*HandshakeMatchSNI)(nil)
)
