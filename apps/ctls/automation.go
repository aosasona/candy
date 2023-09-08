package ctls

import (
	"github.com/aosasona/fate/pkg/candy/apps/ctls/issuers"
	"github.com/aosasona/fate/pkg/candy/storage"
)

type Issuers []issuers.Issuer

type OSCPOverrides map[string]string

type Policy struct {
	Subjects            []string         `json:"subjects"`
	Issuers             Issuers          `json:"issuers,omitempty"`
	GetCertificate      []GetCertificate `json:"get_certificate,omitempty"`
	MustStaple          bool             `json:"must_staple,omitempty"`
	RenewalWindowRatio  int              `json:"renewal_window_ratio,omitempty"`
	KeyType             string           `json:"key_type,omitempty"`
	Storage             storage.Storage  `json:"storage,omitempty"`
	OnDemand            bool             `json:"on_demand,omitempty"`
	DisableOCSPStapling bool             `json:"disable_ocsp_stapling,omitempty"`
	OCSPOverrides       OSCPOverrides    `json:"ocsp_overrides,omitempty"`
}

type OnDemandRateLimit struct {
	Interval int `json:"interval,omitempty"`
	Burst    int `json:"burst,omitempty"`
}

type OnDemand struct {
	RateLimit *OnDemandRateLimit `json:"rate_limit,omitempty"`
	Ask       string             `json:"ask,omitempty"`
}

type Automation struct {
	Policies             []Policy  `json:"policies,omitempty"`
	OnDemand             *OnDemand `json:"on_demand,omitempty"`
	OCSPInterval         int       `json:"ocsp_interval,omitempty"`
	RenewInterval        int       `json:"renew_interval,omitempty"`
	StorageCleanInterval int       `json:"storage_clean_interval,omitempty"`
}
