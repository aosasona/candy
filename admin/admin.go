package admin

import "github.com/aosasona/fate/pkg/candy/apps/ctls/issuers"

type Config struct {
	Persist bool         `json:"persist,omitempty"`
	Load    ConfigLoader `json:"load,omitempty"`
}

type Identity struct {
	Identifiers []string         `json:"identifiers,omitempty"`
	Issuers     []issuers.Issuer `json:"issuers,omitempty"`
}

type Permission struct {
	Paths   []string `json:"paths,omitempty"`
	Methods []string `json:"methods,omitempty"`
}

type AccessControl struct {
	PublicKeys  []string     `json:"public_keys,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type Remote struct {
	Listen        string          `json:"listen,omitempty"`
	AccessControl []AccessControl `json:"access_control,omitempty"`
}

type Admin struct {
	Disabled      bool      `json:"disabled,omitempty"`
	Listen        string    `json:"listen,omitempty"`
	EnforceOrigin bool      `json:"enforce_origin,omitempty"`
	Origins       []string  `json:"origins,omitempty"`
	Config        *Config   `json:"config,omitempty"`
	Identity      *Identity `json:"identity,omitempty"`
	Remote        *Remote   `json:"remote,omitempty"`
}
