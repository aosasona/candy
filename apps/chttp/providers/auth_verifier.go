package providers

type BaseAuthVerifier struct {
	VerifierName AuthVerifierName `json:"verifier"`
}

func (b *BaseAuthVerifier) Verifier(name AuthVerifierName) AuthVerifierName {
	b.VerifierName = name
	return b.VerifierName
}

type AuthVerifierLeaf struct {
	BaseAuthVerifier
}

func NewAuthVerifier[AV AuthVerifier](verifier AV) AV {
	verifier.Verifier()
	return verifier
}

func (av *AuthVerifierLeaf) Verifier() AuthVerifierName {
	return av.BaseAuthVerifier.Verifier(AuthVerifierNameLeaf)
}

// Interface guards
var (
	_ AuthVerifier = (*AuthVerifierLeaf)(nil)
	_ AuthVerifier = NewAuthVerifier(&AuthVerifierLeaf{})
)
