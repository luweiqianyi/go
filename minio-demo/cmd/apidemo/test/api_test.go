package test

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

const (
	endpoint        = "localhost:9000"
	accessKeyID     = "pvT04yBcYxhrXS3zMIT3"
	secretAccessKey = "Cu1c3Wy4EJt6OsJGaCPIQ6eGg8RQx5jpdFtks8ja"
	useSSL          = false
)

func TestConnect(t *testing.T) {
	minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		},
	)

	if err != nil {
		fmt.Printf("connect to remote minio server failed, err: %v", err)
	}

	fmt.Printf("connect to remote minio server success, client info: %v", minioClient)
}

func TestMakeBucket(t *testing.T) {
	ctx := context.Background()

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "music"
	location := "cn-zj-1" // 中国浙江一区

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func TestUploadZip(t *testing.T) {
	ctx := context.Background()

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "music"

	// Upload the zip file
	objectName := "test.zip"
	filePath := "./test.zip"
	contentType := "application/zip"

	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %#v\n", objectName, info)
}

//http://127.0.0.1:9000/music/test.zip?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=XD0XUMWEDQCUFMBKVRHH%2F20230912%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230912T073448Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJYRDBYVU1XRURRQ1VGTUJLVlJISCIsImV4cCI6MTY5NDU0NzAzNCwicGFyZW50IjoiYWRtaW4ifQ.9W8GIi_6ufh0h8D661ZjCZo8WGxaZ-aBUwJtzKZWR8sJVgHkjpweSFczuJo-6-C7iG5zTP81IIv5kia4Klxf4w&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=8f1f84520a2dd5e67a83089bf33e77193df8b26b1943b7250306067a5fac42a0
