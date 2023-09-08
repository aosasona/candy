package ctls

import "github.com/aosasona/fate/pkg/candy/storage"

type KeySourceName string

const (
	KeySourceNameDistributed KeySourceName = "distributed"
	KeySourceNameStandard    KeySourceName = "standard"
)

type KeySource interface {
	Provider() KeySourceName
}

type BaseKeySource struct {
	ProviderName KeySourceName `json:"provider,omitempty"`
}

func (b *BaseKeySource) Provider(name KeySourceName) KeySourceName {
	b.ProviderName = name
	return b.ProviderName
}

func NewkeySource[K KeySource](keySource K) K {
	keySource.Provider()
	return keySource
}

type KeySourceDistributed struct {
	BaseKeySource
	Storage storage.Storage `json:"storage,omitempty"`
}

type KeySourceStandard struct {
	BaseKeySource
}

func (k *KeySourceDistributed) Provider() KeySourceName {
	return k.BaseKeySource.Provider(KeySourceNameDistributed)
}

func (k *KeySourceStandard) Provider() KeySourceName {
	return k.BaseKeySource.Provider(KeySourceNameStandard)
}

// interface guards
var (
	_ KeySource = NewkeySource(&KeySourceDistributed{})
	_ KeySource = NewkeySource(&KeySourceStandard{})

	_ KeySource = (*KeySourceDistributed)(nil)
	_ KeySource = (*KeySourceStandard)(nil)
)
