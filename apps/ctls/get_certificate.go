package ctls

type GetCerfiticateName string

const (
	GetCerfiticateNameHTTP      GetCerfiticateName = "http"
	GetCerfiticateNameTailscale GetCerfiticateName = "tailscale"
)

type GetCertificate interface {
	Via() GetCerfiticateName
}

type BaseGetCertificate struct {
	Name GetCerfiticateName `json:"via"`
}

func (b *BaseGetCertificate) Via(name GetCerfiticateName) GetCerfiticateName {
	b.Name = name
	return b.Name
}

func NewGetCertificate[G GetCertificate](gc G) G {
	gc.Via()
	return gc
}

type GetCertificateHTTP struct {
	BaseGetCertificate
	URL string `json:"url"`
}

type GetCertificateTailscale struct {
	BaseGetCertificate
}

func (g *GetCertificateHTTP) Via() GetCerfiticateName {
	return g.BaseGetCertificate.Via(GetCerfiticateNameHTTP)
}

func (g *GetCertificateTailscale) Via() GetCerfiticateName {
	return g.BaseGetCertificate.Via(GetCerfiticateNameTailscale)
}

// Interface guards
var (
	_ GetCertificate = NewGetCertificate(&GetCertificateHTTP{})
	_ GetCertificate = NewGetCertificate(&GetCertificateTailscale{})

	_ GetCertificate = (*GetCertificateHTTP)(nil)
	_ GetCertificate = (*GetCertificateTailscale)(nil)
)
