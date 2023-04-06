package main

import (
	"github.com/atadzan/file-uploader/pkg/handlers"
	"github.com/atadzan/file-uploader/storage"
	"log"
	"net/http"
)

func main() {
	minioClient, err := storage.NewMinioClient()
	if err != nil {
		log.Fatalf("Error while initializing minio client. Error: %v\n", err.Error())
	}
	h := handlers.NewHandler(minioClient)
	http.HandleFunc("/bucket", h.CreateBucket)
	http.HandleFunc("/buckets", h.GetBuckets)
	http.HandleFunc("/bucket/remove", h.RemoveBucket)
	http.HandleFunc("/upload", h.UploadFile)
	http.HandleFunc("/download", h.DownloadFile)
	http.HandleFunc("/file", h.GetFile)
	http.HandleFunc("/file/remove", h.RemoveFile)

	if err = http.ListenAndServe("localhost:8002", nil); err != nil {
		log.Fatalf("Error while initializing app. Error: %v \n", err)
	}

}
