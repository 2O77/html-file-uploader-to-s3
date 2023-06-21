package storagecontroller

import (
	"fmt"

	storage "github.com/2o77/wope_case/API/internal/domain/storage"
	"github.com/gofiber/fiber/v2"
)

type StorageController struct {
	storageService storage.HTMLStorageService
}

func NewStorageController(storageService storage.HTMLStorageService) (*StorageController, error) {
	storageController := &StorageController{storageService: storageService}

	return storageController, nil
}

func (storageController *StorageController) UploadFile(c *fiber.Ctx) (storage.FileUploadResult, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return storage.FileUploadResult{}, fmt.Errorf("failed to retrieve file from request: %v", err)
	}

	src, err := file.Open()
	if err != nil {
		return storage.FileUploadResult{}, fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	fileContent := make([]byte, file.Size)
	_, err = src.Read(fileContent)
	if err != nil {
		return storage.FileUploadResult{}, fmt.Errorf("failed to read file: %v", err)
	}

	fileUploadResult, err := storageController.storageService.UploadFile(fileContent)
	if err != nil {
		return storage.FileUploadResult{}, fmt.Errorf("failed to upload file: %v", err)
	}

	return fileUploadResult, nil
}
