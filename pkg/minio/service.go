package minio

import (
	"bytes"
	"io"
	"net/url"
	"os"
	"strconv"
	"time"

	"go-base-structure/pkg/file"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var PublicMinioClient *minio.Client
var bucketName string

func Init() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	region := os.Getenv("MINIO_REGION")
	bucketName = os.Getenv("MINIO_BUCKET")
	ssl, _ := strconv.ParseBool(os.Getenv("MINIO_SSL"))

	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
		Region: region,
	})
	if err != nil {
		panic(err)
	}
}

func Upload(c echo.Context, dir string, fileName string, file []byte) error {
	object := bytes.NewReader(file)
	objectSize := int64(len(file))

	_, err := MinioClient.PutObject(
		c.Request().Context(),
		bucketName,
		dir+"/"+fileName,
		object,
		objectSize,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func Remove(c echo.Context, fileName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := MinioClient.RemoveObject(
		c.Request().Context(),
		bucketName,
		fileName,
		opts,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetURL(c echo.Context, dir string, fileName string) (string, error) {
	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-type", "image/png")

	// Gernerate presigned get object url.
	presignedURL, err := MinioClient.PresignedGetObject(
		c.Request().Context(),
		bucketName,
		dir+"/"+fileName,
		time.Duration(1000)*time.Second,
		reqParams,
	)
	if err != nil {
		return "", nil
	}

	return presignedURL.String(), nil
}

func GetObject(c echo.Context, fileName string, dir string) (string, error) {
	var destination string
	reader, err := MinioClient.GetObject(
		c.Request().Context(),
		bucketName,
		dir+fileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return destination, err
	}
	defer reader.Close()

	rootPath := file.GetRootDirectory()
	destination = rootPath + "/public/" + fileName
	localFile, err := os.Create(destination)
	if err != nil {
		return destination, err
	}
	defer localFile.Close()

	stat, err := reader.Stat()
	if err != nil {
		return destination, err
	}

	if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
		return destination, err
	}

	return destination, nil
}
