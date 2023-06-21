package storage

type FileUploadResult struct {
	FilePath string
}

type HTMLStorageRepository interface {
	UploadFile(data []byte) (FileUploadResult, error)
}

type HTMLStorageService interface {
	UploadFile(data []byte) (FileUploadResult, error)
}
