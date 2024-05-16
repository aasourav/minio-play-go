package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "localhost:9033"
	accessKeyID := "SVlNTG1fqbmsq7B8qEGk"
	secretAccessKey := "OUYglwDq5CUmt1lP0YbpY1RPbCxS5ST8Ak8N4j8k"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// log.Printf("%#v\n", minioClient) // minioClient is now setup
	// err = minioClient.MakeBucket(context.Background(), "new-bucket", minio.MakeBucketOptions{Region: "api-south"})

	// if err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Printf("success")
	// }

	// err = minioClient.FGetObject(context.Background(), "aas", "bleachbit_4.6.0-0_all_ubuntu2204.deb", "downloads.deb", minio.GetObjectOptions{})

	file, err := os.Open("./cookies.txt")
	defer file.Close()

	fileStat, err := file.Stat()
	minioClient.PutObject(context.Background(), "aas", "test-obj.txt", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+"test-obj.txt"+"\"")
	url, _ := minioClient.PresignedGetObject(context.Background(), "aas", "test-obj.txt", time.Hour*24*7, reqParams)
	log.Println(url)
}
