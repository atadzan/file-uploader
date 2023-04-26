package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func NewMinioClient() (*minio.Client, error) {
	endpoint := "192.168.1.61:9099"
	//accessKeyId := "MKJdVEufsJngEP8D"
	accessKeyId := "fxCbzOL1BiTYb2f1"
	//accessKeyId := "et2MPdWYdNPyigBU"
	//accessKeyId := "rtHvlzo9Sbhpg3bb"
	//secretAccessKey := "1CyMiAdvJFBEgWLyAZM0ukQNesctJbe5"
	secretAccessKey := "BXpAbcwRp7dTM1IVijqYD3ojbl2jMijc"
	//secretAccessKey := "SAm0PBUOxzv0dqmzu65gEd6vJFnJxS0o"
	//secretAccessKey := "hgzfhn3DqlaVs03QFAqRqumrDMurZTEk"

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
	ctx := context.Background()
	status, err := minioClient.BucketExists(ctx, "videos")
	if err != nil {
		return nil, fmt.Errorf("failed to check minio bucket. Error %v", err)
	}
	if status != true {
		err = minioClient.MakeBucket(ctx, "videos", minio.MakeBucketOptions{})
	}
	return minioClient, nil
}
