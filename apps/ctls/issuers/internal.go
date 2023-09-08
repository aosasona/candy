package issuers

type IssuerInternal struct {
	BaseIssuer
	CA           string `json:"ca,omitempty"`
	Lifetime     int    `json:"lifetime,omitempty"`
	SignWithRoot bool   `json:"sign_with_root,omitempty"`
}

func (i *IssuerInternal) Name() IssuerName { return i.BaseIssuer.Name(IssuerNameInternal) }
