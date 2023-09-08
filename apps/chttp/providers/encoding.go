package providers

type (
	EncodingZstd struct{}

	EncodingGzip struct {
		Level int `json:"level"`
	}
)

func (e EncodingGzip) Encoding() EncodingName { return EncodingNameGzip }

func (e EncodingZstd) Encoding() EncodingName { return EncodingNameZstd }
