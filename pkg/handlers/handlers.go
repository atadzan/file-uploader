package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/atadzan/file-uploader/models"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Handler struct {
	storage *minio.Client
}

func NewHandler(minio *minio.Client) *Handler {
	return &Handler{
		storage: minio,
	}
}

func (h *Handler) CreateBucket(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	var input models.BucketInput
	if err = json.Unmarshal(body, &input); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	err = h.storage.MakeBucket(r.Context(), input.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := h.storage.BucketExists(r.Context(), input.BucketName)
		if errBucketExists == nil && exists {
			GenerateResponse(models.ResponseParam{
				W:       w,
				Message: "Bucket exists",
				Status:  http.StatusBadRequest,
			})
			return
		}
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: "Successfully created",
		Status:  http.StatusOK,
	})
	return
}

func (h *Handler) GetBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := h.storage.ListBuckets(r.Context())
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	response, err := json.Marshal(buckets)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	w.Write(response)
	w.WriteHeader(http.StatusOK)
	return
}

func (h *Handler) RemoveBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := r.URL.Query().Get("title")
	if err := h.storage.RemoveBucket(r.Context(), bucketName); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: "Successfully removed",
		Status:  http.StatusOK,
	})
	return
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	var input models.UploadInput
	if err = json.Unmarshal(body, &input); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	info, err := h.storage.FPutObject(r.Context(), input.BucketName, input.FileName, input.FilePath, minio.PutObjectOptions{ContentType: input.ContentType})
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: fmt.Sprintf("Successfully uploaded %s of size %d \n", input.FileName, info.Size),
		Status:  http.StatusOK,
	})
	return
}

func (h *Handler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	var input models.DownloadInput
	if err = json.Unmarshal(body, &input); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	object, err := h.storage.GetObject(r.Context(), input.BucketName, input.FileName, minio.GetObjectOptions{})
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer object.Close()

	localFile, err := os.Create(input.DestinationPath)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: fmt.Sprintf("Successfully downloaded in %s", input.DestinationPath),
		Status:  http.StatusOK,
	})
	return
}

func (h *Handler) RemoveFile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	var input models.GetFile
	if err = json.Unmarshal(body, &input); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	err = h.storage.RemoveObject(r.Context(), input.BucketName, input.FileName, minio.RemoveObjectOptions{})
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: "Successfully removed",
		Status:  http.StatusOK,
	})
	return
}

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	var input models.GetFile
	if err = json.Unmarshal(body, &input); err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	reqParam := make(url.Values)
	reqParam.Set("response-content-disposition", "attachment; filename=\""+input.FileName+"\"")
	presignedUrl, err := h.storage.PresignedGetObject(r.Context(), input.BucketName, input.FileName, time.Second*60, reqParam)
	if err != nil {
		GenerateResponse(models.ResponseParam{
			W:       w,
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	GenerateResponse(models.ResponseParam{
		W:       w,
		Message: presignedUrl.String(),
		Status:  http.StatusOK,
	})
	return
}

func GenerateResponse(param models.ResponseParam) {
	rawResponse := models.Response{
		Code:    param.Status,
		Message: param.Message,
	}
	response, err := json.Marshal(rawResponse)
	if err != nil {
		log.Println(err.Error())
	}
	param.W.Write(response)
	param.W.WriteHeader(param.Status)
	return
}
