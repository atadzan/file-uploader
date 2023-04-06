package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func NewMinioClient() (*minio.Client, error) {
	endpoint := "192.168.1.61:9000"
	accessKeyId := "MKJdVEufsJngEP8D"
	secretAccessKey := "1CyMiAdvJFBEgWLyAZM0ukQNesctJbe5"
	useSSL := false

	//Initialize minIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err.Error())
		return minioClient, err
	}
	return minioClient, nil
}
