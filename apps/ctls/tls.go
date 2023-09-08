package ctls

type Certificates map[CertLoaderName]CertLoader

type SessionTickets struct {
	KeySource        KeySource `json:"key_source,omitempty"`
	RotationInterval int       `json:"rotation_interval,omitempty"`
	MaxKeys          int       `json:"max_keys,omitempty"`
	DisableRotation  bool      `json:"disable_rotation,omitempty"`
	Disabled         bool      `json:"disabled,omitempty"`
}

type Cache struct {
	Capacity int `json:"capacity,omitempty"`
}

type TLS struct {
	Certificates        *Certificates   `json:"certificates,omitempty"`
	Automation          *Automation     `json:"automation,omitempty"`
	SessionTickets      *SessionTickets `json:"session_tickets,omitempty"`
	Cache               *Cache          `json:"cache,omitempty"`
	DisableOCSPStapling bool            `json:"disable_ocsp_stapling,omitempty"`
}
