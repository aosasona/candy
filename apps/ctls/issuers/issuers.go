package issuers

type IssuerName string

const (
	IssuerNameACME     IssuerName = "acme"
	IssuerNameInternal IssuerName = "internal"
	IssuerNameZeroSSL  IssuerName = "zerossl"
)

type Issuer interface {
	Name() IssuerName
}

type BaseIssuer struct {
	IssuerName IssuerName `json:"module"`
}

func (b *BaseIssuer) Name(name IssuerName) IssuerName {
	b.IssuerName = name
	return b.IssuerName
}

func NewIssuer[I Issuer](issuer I) I {
	issuer.Name()
	return issuer
}

// Interface guards
var (
	_ Issuer = NewIssuer(&IssuerAcme{})
	_ Issuer = NewIssuer(&IssuerInternal{})
	_ Issuer = NewIssuer(&IssuerZeroSSL{})

	_ Issuer = (*IssuerAcme)(nil)
	_ Issuer = (*IssuerInternal)(nil)
	_ Issuer = (*IssuerZeroSSL)(nil)
)
