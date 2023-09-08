package providers

type (
	AuthProviderName string
	EncodingName     string
	AuthVerifierName string
	TrustedProxyName string
)

const (
	AuthProviderNameHTTPBasic AuthProviderName = "http_basic"

	AuthVerifierNameLeaf AuthVerifierName = "leaf"

	EncodingNameGzip EncodingName = "gzip"
	EncodingNameZstd EncodingName = "zstd"

	TrustedProxyNameStatic TrustedProxyName = "static"
)

type AuthProvider interface {
	Provider() AuthProviderName
}

type Encoding interface {
	Encoding() EncodingName
}

type AuthVerifier interface {
	Verifier() AuthVerifierName
}

type TrustedProxy interface {
	Source() TrustedProxyName
}
