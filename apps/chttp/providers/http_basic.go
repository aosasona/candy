package providers

// Algorithms definitions

type (
	Algorithm string

	Hash interface {
		SetAlgorithm()
	}

	BaseHash struct {
		Algorithm Algorithm `json:"algorithm"`
	}
)

const (
	AlgorithmBcrypt Algorithm = "bcrypt"
	AlgorithmScrypt Algorithm = "scrypt"
)

type HashBcrypt struct {
	BaseHash
}

type HashScrypt struct {
	BaseHash
	N         int `json:"N,omitempty"`
	R         int `json:"r,omitempty"`
	P         int `json:"p,omitempty"`
	KeyLength int `json:"key_length,omitempty"`
}

func NewHash[H Hash](hash H) H {
	hash.SetAlgorithm()
	return hash
}

func (h *HashBcrypt) SetAlgorithm() {
	h.Algorithm = AlgorithmBcrypt
}

func (h *HashScrypt) SetAlgorithm() {
	h.Algorithm = AlgorithmScrypt
}

// Provider definitions
type AuthAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt,omitempty"`
}

type AuthProviderHTTPBasic struct {
	Hash     Hash          `json:"hash,omitempty"`
	Accounts []AuthAccount `json:"accounts,omitempty"`
	Realm    string        `json:"realm,omitempty"`
	HashCache map[string]string `json:"hash_cache,omitempty"`
}

func (p AuthProviderHTTPBasic) Provider() AuthProviderName {
	return AuthProviderNameHTTPBasic
}

// Interface guards
var (
	_ Hash = NewHash(&HashBcrypt{})

	_ AuthProvider = (*AuthProviderHTTPBasic)(nil)
	_ Hash         = (*HashBcrypt)(nil)
	_ Hash         = (*HashScrypt)(nil)
)
