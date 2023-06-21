package storagerepository

import (
	"bytes"
	"fmt"
	"os"

	storage "github.com/2o77/wope_case/API/internal/domain/storage"
	idGenerator "github.com/2o77/wope_case/API/internal/generator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	godotenv "github.com/joho/godotenv"
)

type S3Repository struct {
	s3Client *s3.S3
}

var awsRegion string
var awsBucket string
var awsAccessKeyID string
var awsSecretAccessKey string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	awsRegion = os.Getenv("AWS_REGION")
	awsBucket = os.Getenv("AWS_BUCKET")
	awsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func NewS3Repository() (*S3Repository, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil, err
	}

	s3Client := s3.New(sess)

	return &S3Repository{
		s3Client: s3Client,
	}, nil
}

func (s3Repo *S3Repository) UploadFile(data []byte) (storage.FileUploadResult, error) {
	fileName, err := idGenerator.NewIDGenerator()
	if err != nil {
		fmt.Println("Error generating random ID:", err)
		return storage.FileUploadResult{FilePath: ""}, err
	}

	_, err = s3Repo.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(awsBucket),
		Key:         aws.String("html-files/" + fileName + ".html"),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("multipart/form-data"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return storage.FileUploadResult{FilePath: ""}, err
	}

	return storage.FileUploadResult{FilePath: "https://" + awsBucket + ".s3.amazonaws.com/index.html"}, nil
}
