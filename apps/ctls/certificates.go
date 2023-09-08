package ctls

type CertLoaderName string

const (
	CertLoaderNameAutomate CertLoaderName = "automate"
	CertLoaderNameFiles    CertLoaderName = "load_files"
	CertLoaderNameFolders  CertLoaderName = "load_folders"
	CertLoaderNamePem      CertLoaderName = "load_pem"
	CertLoaderNameStorage  CertLoaderName = "load_storage"
)

type CertLoader interface {
	Name() CertLoaderName
}

type CertLoaderAutomate []string

type (
	FileLoader struct {
		Certificate string   `json:"certificate"`
		Key         string   `json:"key"`
		Format      string   `json:"format"`
		Tags        []string `json:"tags"`
	}

	CertLoaderFiles []FileLoader
)

type CertLoaderFolders []string

type (
	PemLoader struct {
		Certificate string   `json:"certificate"`
		Key         string   `json:"key"`
		Tags        []string `json:"tags"`
	}
	CertLoaderPem []PemLoader
)

type CertLoaderStorage struct {
	Pairs FileLoader `json:"pairs"`
}

// Methods
func (c CertLoaderAutomate) Name() CertLoaderName { return CertLoaderNameAutomate }

func (c CertLoaderFiles) Name() CertLoaderName { return CertLoaderNameFiles }

func (c CertLoaderFolders) Name() CertLoaderName { return CertLoaderNameFolders }

func (c CertLoaderPem) Name() CertLoaderName { return CertLoaderNamePem }

func (c CertLoaderStorage) Name() CertLoaderName { return CertLoaderNameStorage }

// Interface guards
var (
	_ CertLoader = CertLoaderAutomate{}
	_ CertLoader = CertLoaderFiles{}
	_ CertLoader = CertLoaderFolders{}
	_ CertLoader = CertLoaderPem{}
	_ CertLoader = CertLoaderStorage{}
)
