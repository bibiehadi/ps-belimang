package uploadcontroller

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (controller *uploadController) UploadImage(c echo.Context) error {

	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalid file",
		})
	}

	if file.Size > int64(2*1024*1024) {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Image is too large, maximum size is 2MB",
		})
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Error Opening  File",
		})
	}
	defer src.Close()

	if !isImages(file) {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "File is not image, require file is jpg or jpeg image",
		})
	}

	//rename file name to uuid

	fileExt := filepath.Ext(file.Filename)
	uuid := uuid.New().String()
	key := fmt.Sprintf("upload/%s%s", uuid, fileExt)

	//upload to S3
	bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	_, err = helpers.S3Client().PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   src,
	})

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed uploaded image to S3 ",
		})
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, key)

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Success uploaded image",
		Data: entities.ImageUploadResponse{
			URL: fileURL,
		},
	})
}

func isImages(file *multipart.FileHeader) bool {
	src, err := file.Open()
	if err != nil {
		return false
	}
	defer src.Close()

	buf := make([]byte, 512)
	_, err = src.Read(buf)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buf)
	return strings.HasPrefix(contentType, "image/jpeg")
}
