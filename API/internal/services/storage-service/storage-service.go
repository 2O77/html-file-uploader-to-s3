package storageservice

import (
	"fmt"

	storage "github.com/2o77/wope_case/API/internal/domain/storage"
)

type StorageService struct {
	storageRepository storage.HTMLStorageRepository
}

func NewStorageService(storageRepository storage.HTMLStorageRepository) (*StorageService, error) {
	storageService := &StorageService{storageRepository: storageRepository}

	return storageService, nil
}

func (storageService *StorageService) UploadFile(fileContent []byte) (storage.FileUploadResult, error) {
	fileUploadResult, err := storageService.storageRepository.UploadFile(fileContent)
	if err != nil {
		return storage.FileUploadResult{}, fmt.Errorf("failed to upload file: %v", err)
	}

	return fileUploadResult, nil
}
