package storage

type StorageName string

const (
	StorageNameFileSystem StorageName = "file_system"
)

type Storage interface {
	Name() StorageName
}

type BaseStorage struct {
	StorageName StorageName `json:"module"`
}

func (b *BaseStorage) Name(name StorageName) StorageName {
	b.StorageName = name
	return b.StorageName
}

func NewStorage[S Storage](s S) S {
	s.Name()
	return s
}

type FileSystemStorage struct {
	BaseStorage
	Root string `json:"root"`
}

func (s *FileSystemStorage) Name() StorageName {
	return s.BaseStorage.Name(StorageNameFileSystem)
}

// Interface guards
var (
	_ Storage = NewStorage(&FileSystemStorage{})

	_ Storage = (*FileSystemStorage)(nil)
)
