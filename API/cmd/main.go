package main

import (
	"fmt"
	"path/filepath"

	storagecontroller "github.com/2o77/wope_case/API/internal/controllers/storage-controller"
	messagequeue "github.com/2o77/wope_case/API/internal/domain/message-queue"
	kafkastorage "github.com/2o77/wope_case/API/internal/message-queue/kafka"
	storagerepository "github.com/2o77/wope_case/API/internal/repositories/storage-repository"
	storageservice "github.com/2o77/wope_case/API/internal/services/storage-service"

	fiber "github.com/gofiber/fiber/v2"
)

type StorageController struct {
	storagecontroller.StorageController
}

func main() {

	storageRepository, err := storagerepository.NewS3Repository()
	if err != nil {
		fmt.Println(err)
	}
	storageService, err := storageservice.NewStorageService(storageRepository)
	if err != nil {
		fmt.Println(err)
	}

	storageController, err := storagecontroller.NewStorageController(storageService)
	if err != nil {
		fmt.Println(err)
	}

	kafkaStorage, err := kafkastorage.NewKafka()
	if err != nil {
		fmt.Println(err)
	}

	// err = kafkaStorage.PublishMessage(messagequeue.Message{FilePath: "akdfjasdlkfasd"})
	// if err != nil {
	// 	panic(err)
	// }

	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		path, err := storageController.UploadFile(c)
		if err != nil {
			return err
		}

		filePath := filepath.Base(path.FilePath)

		err = kafkaStorage.PublishMessage(messagequeue.Message{FilePath: filePath})
		if err != nil {
			return err
		}

		return c.Status(200).SendString("File uploaded successfully")
	})

	app.Listen(":3000")
}
