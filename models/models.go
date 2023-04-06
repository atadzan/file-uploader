package models

import "net/http"

type BucketInput struct {
	BucketName string `json:"bucketName"`
}

type UploadInput struct {
	BucketName  string `json:"bucketName"`
	FileName    string `json:"fileName"`
	FilePath    string `json:"filePath"`
	ContentType string `json:"contentType"`
}

type DownloadInput struct {
	BucketName      string `json:"bucketName"`
	FileName        string `json:"fileName"`
	DestinationPath string `json:"destinationPath"`
}

type GetFile struct {
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
}

type ResponseParam struct {
	W       http.ResponseWriter
	Message interface{}
	Status  int
}

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
