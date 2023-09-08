package providers

type BaseTrustedProxy struct {
	SourceName TrustedProxyName `json:"source"`
}

func (t *BaseTrustedProxy) Source(name TrustedProxyName) TrustedProxyName {
	t.SourceName = name
	return t.SourceName
}

func NewTrustedProxy[T TrustedProxy](proxy T) T {
	proxy.Source()
	return proxy
}

type TrustedProxyStatic struct {
	BaseTrustedProxy
	Ranges []string `json:"ranges"`
}

func (t *TrustedProxyStatic) Source() TrustedProxyName {
	return t.BaseTrustedProxy.Source(TrustedProxyNameStatic)
}

// Intetface guards
var (
	_ TrustedProxy = NewTrustedProxy(&TrustedProxyStatic{})

	_ TrustedProxy = (*TrustedProxyStatic)(nil)
)
